package handler

import (
	"fmt"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/config/logger"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/service"
	"github.com/gofiber/fiber/v2"
)

var log = logger.NewLogger()

type UserHandler struct {
	s service.Service
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	payload := new(model.UserPayload)

	if err := ctx.BodyParser(payload); err != nil {
		log.WithError(err).Errorf("bad request error, method: %v, path: %v", ctx.Method(), ctx.Path())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		errorValidate := fmt.Errorf("validate error")
		log.WithError(errorValidate).Errorf("bad request error, method: %v, path: %v", ctx.Method(), ctx.Path())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.s.User.InsertUser(ctx.Context(), *payload); err != nil {
		log.WithError(err).Errorf("intenal server error, method: %v, path: %v", ctx.Method(), ctx.Path())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
	})
}
