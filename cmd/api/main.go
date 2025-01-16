package main

import (
	"fmt"
	"go-rest-api-kafka/internal/config"
	"go-rest-api-kafka/internal/database"
	"go-rest-api-kafka/internal/server"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к базе данных с повторными попытками
	var db *gorm.DB
	maxRetries := 5
	retryDelay := time.Second * 5

	for i := 0; i < maxRetries; i++ {
		db, err = database.NewPostgresDB(cfg)
		if err == nil {
			break
		}

		if i == maxRetries-1 {
			log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
		}

		// Проверяем, является ли ошибка связанной с SSL
		if cfg.Postgres.SSL == "Y" && isSSLError(err) {
			log.Fatalf("SSL configuration error: %v", err)
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	// Проверка соединения с базой данных
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Установка параметров пула соединений
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	// Проверка живости соединения
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	// Создание и запуск HTTP сервера
	srv := server.NewServer(db)
	if err := srv.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Printf("Server shutdown error: %v", err)
		os.Exit(1)
	}
}

// isSSLError проверяет, связана ли ошибка с SSL-конфигурацией
func isSSLError(err error) bool {
	errMsg := err.Error()
	sslErrors := []string{
		"certificate",
		"SSL",
		"sslmode",
		"x509",
		"tls",
	}

	for _, sslErr := range sslErrors {
		if strings.Contains(strings.ToLower(errMsg), strings.ToLower(sslErr)) {
			return true
		}
	}
	return false
}