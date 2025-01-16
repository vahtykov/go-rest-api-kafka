package database

import (
	"fmt"
	"go-rest-api-kafka/internal/config"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	var dsn string
	
	if cfg.Postgres.SSL == "Y" {
		// Проверяем наличие SSL сертификатов
		certFiles := map[string]string{
			"/vault/secrets/tls_root_pem.txt": "sslrootcert",
			"/vault/secrets/tls.key":          "sslkey",
			"/vault/secrets/tls_pem.txt":      "sslcert",
		}

		// Проверяем существование всех необходимых файлов сертификатов
		for file := range certFiles {
			if _, err := os.Stat(file); err != nil {
				return nil, fmt.Errorf("SSL certificate file not found: %s", file)
			}
		}

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s "+
			"sslmode=verify-full "+
			"sslrootcert=/vault/secrets/tls_root_pem.txt "+
			"sslkey=/vault/secrets/tls.key "+
			"sslcert=/vault/secrets/tls_pem.txt",
			cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, 
			cfg.Postgres.Database, cfg.Postgres.Port,
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, 
			cfg.Postgres.Database, cfg.Postgres.Port,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.Postgres.Schema + ".",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}