package database

import (
	"context"
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Setup(ctx context.Context) (*pgxpool.Pool, error) {
	db, err := Connect(ctx, pgxpool.New)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := RunMigrations(db); err != nil {
		return nil, err
	}
	return db, nil
}

func Name() (string, error) {
	value, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		return "", fmt.Errorf("DATABASE_NAME environment variable not set")
	}
	return value, nil
}

func URL() (string, error) {
	value, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return "", fmt.Errorf("DATABASE_URL environment variable not set")
	}
	return value, nil
}

//go:embed schema/*.sql
var migrations embed.FS

func RunMigrations(pool *pgxpool.Pool) error {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect("pgx"); err != nil {
		return fmt.Errorf("could not set goose dialect: %w", err)
	}
	db := stdlib.OpenDBFromPool(pool)
	db.SetMaxIdleConns(0)
	if err := goose.Up(db, "schema"); err != nil {
		return fmt.Errorf("could not run migrations: %w", err)
	}
	return nil
}

const ConnectPingTimeout = 5 * time.Second

type Pinger interface {
	DBTX
	Ping(context.Context) error
}

type NewFunc[DB Pinger] func(ctx context.Context, dbURL string) (DB, error)

func Connect[DB Pinger](ctx context.Context, newDB NewFunc[DB]) (DB, error) {
	databaseURL, err := URL()
	if err != nil {
		var zero DB
		return zero, fmt.Errorf("could not determine database URL: %w", err)
	}
	dbConn, err := newDB(ctx, databaseURL)
	if err != nil {
		var zero DB
		return zero, fmt.Errorf("failed to create database pool: %w", err)
	}
	pingCtx, cancel := context.WithTimeout(ctx, ConnectPingTimeout)
	defer cancel()
	if err := dbConn.Ping(pingCtx); err != nil {
		return dbConn, fmt.Errorf("failed to ping database: %w", err)
	}
	return dbConn, nil
}
