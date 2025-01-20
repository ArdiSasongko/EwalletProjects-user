package handler

import (
	"github.com/ArdiSasongko/EwalletProjects-user/internal/auth"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/service"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Health interface {
		CheckHealth(*fiber.Ctx) error
	}
	User interface {
		Register(*fiber.Ctx) error
		Login(*fiber.Ctx) error
		Logout(*fiber.Ctx) error
		RefreshToken(*fiber.Ctx) error
	}
	Middleware interface {
		AuthMiddleware() fiber.Handler
		RefreshTokenMiddleware() fiber.Handler
	}
}

func NewHandler(db sqlc.DBTX, auth auth.Authenticator) Handlers {
	q := sqlc.New(db)
	service := service.NewService(db, auth)
	return Handlers{
		Health: &HealthHandler{},
		User: &UserHandler{
			s: service,
		},
		Middleware: &MiddlewareHandler{
			service: service,
			q:       q,
			auth:    auth,
		},
	}
}
