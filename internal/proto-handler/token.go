package protohandler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/auth"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/proto/token"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	token.UnimplementedTokenServiceServer
	auth auth.Authenticator
	db   *sqlc.Queries
}

func NewTokenService(db *sqlc.Queries, auth auth.Authenticator) *TokenService {
	return &TokenService{
		auth: auth,
		db:   db,
	}
}

func (s *TokenService) Validate(ctx context.Context, req *token.TokenRequest) (*token.TokenResponse, error) {
	if req.Token == "" {
		errs := errors.New("token is required")
		return &token.TokenResponse{
			Message: errs.Error(),
		}, nil
	}

	// check in database
	_, err := s.db.GetTokenByToken(ctx, req.Token)
	if err != nil {
		return &token.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	// extract token
	jwtToken, err := s.auth.ValidateToken(req.Token)
	if err != nil {
		return &token.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)

	userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
	if err != nil {
		return &token.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &token.TokenResponse{
		Message: "Token Valid",
		Data: &token.UserData{
			Id: int32(userID),
		},
	}, nil
}
