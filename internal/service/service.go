package service

import (
	"context"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/auth"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/external/wallet"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
)

type Service struct {
	User interface {
		InsertUser(context.Context, model.UserPayload) (wallet.WalletResponse, error)
		GetUser(context.Context, model.UserLoginPayload) (*model.LoginResponse, error)
		GetUserByID(context.Context, int32) (*sqlc.User, error)
		DeleteTokenByID(context.Context, int32) error
		RefreshToken(context.Context, *sqlc.User) (*model.LoginResponse, error)
	}
}

func NewService(db sqlc.DBTX, auth auth.Authenticator) Service {
	q := sqlc.New(db)
	walletClient := wallet.NewWalletClient()
	return Service{
		User: &UserService{
			q:      q,
			auth:   auth,
			wallet: walletClient,
		},
	}
}
