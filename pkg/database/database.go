package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/marcelofabianov/my-cash/config"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(conn *sql.DB) *Database {
	return &Database{conn: conn}
}

func (d *Database) Conn() *sql.DB {
	return d.conn
}

func (d *Database) Close() error {
	return d.conn.Close()
}

func (p *Database) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return p.conn.PingContext(ctx)
}

func Connect(ctx context.Context, cfg config.DatabaseConfig) (*Database, error) {
	dsn := FormatDSN(cfg)

	conn, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return NewDatabase(conn), nil
}

func FormatDSN(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
