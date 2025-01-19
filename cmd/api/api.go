package api

import (
	"github.com/ArdiSasongko/EwalletProjects-user/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type application struct {
	handler handler.Handlers
	config  Config
}

type Config struct {
	addr   string
	logger *logrus.Logger
	db     DBConfig
}

type DBConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *fiber.App {
	r := fiber.New()

	v1 := r.Group("/v1")
	v1.Get("/health", app.handler.Health.CheckHealth)

	// authentication handler
	auth := v1.Group("/authentication")
	auth.Post("/register", app.handler.User.Register)
	return r
}

func (app *application) run(r *fiber.App) error {
	app.config.logger.Printf("http server has running, port:%v", app.config.addr)
	return r.Listen(app.config.addr)
}
