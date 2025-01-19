package handler

import (
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
	}
}

func NewHandler(db sqlc.DBTX) Handlers {
	service := service.NewService(db)
	return Handlers{
		Health: &HealthHandler{},
		User: &UserHandler{
			s: service,
		},
	}
}
