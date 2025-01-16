package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres         *PostgresConfig
	KafkaBrokers     string
	KafkaTopic       string
	Port             string
	ServiceConsumerURL string
	DBMaxIdleConns    int
	DBMaxOpenConns    int
	DBConnMaxLifetime time.Duration
}

type PostgresConfig struct {
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSL      string `json:"ssl"`
}

func LoadPostgresConfig(path string) (*PostgresConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading postgres config file: %w", err)
	}

	var config PostgresConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("error parsing postgres config: %w", err)
	}

	return &config, nil
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	pgConfig, err := LoadPostgresConfig("conf/postgres.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load postgres config: %w", err)
	}

	maxIdleConns, err := strconv.Atoi(getEnvOrDefault("DB_MAX_IDLE_CONNS", "10"))
	if err != nil {
		maxIdleConns = 10
	}

	maxOpenConns, err := strconv.Atoi(getEnvOrDefault("DB_MAX_OPEN_CONNS", "100"))
	if err != nil {
		maxOpenConns = 100
	}

	connMaxLifetime, err := time.ParseDuration(getEnvOrDefault("DB_CONN_MAX_LIFETIME", "1h"))
	if err != nil {
		connMaxLifetime = time.Hour
	}

	return &Config{
		Postgres:          pgConfig,
		KafkaBrokers:     os.Getenv("KAFKA_BROKERS"),
		KafkaTopic:       os.Getenv("KAFKA_TOPIC"),
		Port:             os.Getenv("PORT"),
		ServiceConsumerURL: os.Getenv("SERVICE_CONSUMER_URL"),
		DBMaxIdleConns:    maxIdleConns,
		DBMaxOpenConns:    maxOpenConns,
		DBConnMaxLifetime: connMaxLifetime,
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}