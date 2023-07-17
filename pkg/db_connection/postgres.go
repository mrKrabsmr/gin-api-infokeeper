package db_connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrKrabsmr/infokeeper/internal/config"
	"os"
	"time"
)

type PostgresDB struct {
	config *config.Config
}

func NewPGConnection(config *config.Config) *PostgresDB {
	return &PostgresDB{
		config: config,
	}
}

func (p *PostgresDB) PostgreSQLConnection() (*sqlx.DB, error) {
	maxConn := p.config.MaxConn
	maxIdleConn := p.config.MaxIdleConn
	maxLifetimeConn := p.config.MaxLifetimeConn

	db, err := sqlx.Connect("postgres", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
