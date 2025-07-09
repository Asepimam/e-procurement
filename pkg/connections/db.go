package connections

import (
	"database/sql"
	"fmt"
	"time"
)

type DBConfig struct {
	Driver 			string
	DSN 			string
	MaxOpenConns 	int
	MaxIdleConns 	int
	ConnMaxLifetime time.Duration
}
func ConnectDB(cfg DBConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
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