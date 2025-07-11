package connections

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DBConfig struct {
    Driver         string
    Host           string
    Port           string
    User           string
    Password       string
    DBName         string
    Schema         string // tambahan untuk schema
    MaxOpenConns   int
    MaxIdleConns   int
    ConnMaxLifetime time.Duration
}
func ConnectDB(cfg DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
        cfg.Host,
        cfg.Port,
        cfg.User,
        cfg.Password,
        cfg.DBName,
        cfg.Schema,
    )
    log.Printf("Connecting to database with DSN: %s", dsn)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	return db, nil
}