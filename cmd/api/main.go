package main

import (
	"fmt"
	"go-rest-api-kafka/internal/config"
	"go-rest-api-kafka/internal/database"
	"go-rest-api-kafka/internal/server"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к базе данных
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Создание и запуск HTTP сервера
	srv := server.NewServer(db)
	if err := srv.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}