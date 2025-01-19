package api

import (
	"log"

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
		addr:   env.GetEnv("ADDR_HTTP", ":4000"),
		logger: logrus,
		db: DBConfig{
			DB_ADDR: "some address",
		},
	}

	handler := handler.NewHandler()

	app := application{
		config:  cfg,
		handler: handler,
	}

	api := app.mount()
	if err := app.run(api); err != nil {
		cfg.logger.Fatalf("failed to starting http server, err:%v", err)
	}
}
