package model

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type UserPayload struct {
	Username    string `json:"username" validate:"required,min=5,max=255"`
	Email       string `json:"email" validate:"required,email,max=255"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	DoB         string `json:"dob" validate:"required"`
	Password    string `json:"password" validate:"required,min=5,max=255"`
	Fullname    string `json:"fullname" validate:"required,max=255"`
}

func (u *UserPayload) Validate() error {
	return Validate.Struct(u)
}
