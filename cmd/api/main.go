package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/med-IDBENOUAKRIM/lets_go/cmd/utils"
	"github.com/med-IDBENOUAKRIM/lets_go/internal/data"
)

const version = "1.0.0"

type DbConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

type Config struct {
	env string
	db  DbConfig
}

type Application struct {
	config Config
	logger *slog.Logger
	models data.Models
}

func main() {
	var cfg Config

	utils.LoadConfig()
	dbSource := os.Getenv("DB_SOURCE")

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", dbSource, "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connection")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connection")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgeSQL max connection idle time")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("database connection pool established")

	app := &Application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxLifetime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
