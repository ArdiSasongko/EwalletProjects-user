package api

import (
	"context"
	"log"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/config/db"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/config/logger"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/env"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/handler"
	"github.com/joho/godotenv"
)

func SetupHTTP() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// set up logrus
	logrus := logger.NewLogger()

	cfg := Config{
		addr:   env.GetEnvString("ADDR_HTTP", ":4000"),
		logger: logrus,
		db: DBConfig{
			addr:         env.GetEnvString("DB_ADDR", ""),
			maxOpenConns: env.GetEnvInt("DB_MAX_CONNS", 5),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE", 5),
			maxIdleTime:  env.GetEnvString("DB_MAX_TIME_IDLE", "10m"),
		},
	}

	conn, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		cfg.logger.Fatalf("failed to connected database :%v", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		cfg.logger.Fatalf("failed to ping database :%v", err)
	}

	cfg.logger.Info("success connected to database")

	handler := handler.NewHandler(conn)

	app := application{
		config:  cfg,
		handler: handler,
	}

	api := app.mount()
	if err := app.run(api); err != nil {
		cfg.logger.Fatalf("failed to starting http server, err:%v", err)
	}
}
