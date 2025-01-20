package service

import (
	"context"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/auth"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
)

type Service struct {
	User interface {
		InsertUser(context.Context, model.UserPayload) error
		GetUser(context.Context, model.UserLoginPayload) (*model.LoginResponse, error)
		GetUserByID(context.Context, int32) (*sqlc.User, error)
		DeleteTokenByID(context.Context, int32) error
	}
}

func NewService(db sqlc.DBTX, auth auth.Authenticator) Service {
	q := sqlc.New(db)
	return Service{
		User: &UserService{
			q:    q,
			auth: auth,
		},
	}
}
