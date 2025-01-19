package service

import (
	"context"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
)

type Service struct {
	User interface {
		InsertUser(context.Context, model.UserPayload) error
		GetUser(context.Context) error
	}
}

func NewService(db sqlc.DBTX) Service {
	q := sqlc.New(db)
	return Service{
		User: &UserService{
			q: q,
		},
	}
}
