package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/auth"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/external"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/external/wallet"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	q      *sqlc.Queries
	auth   auth.Authenticator
	wallet wallet.WalletClient
	notif  external.Notification
}

func (s *UserService) InsertUser(ctx context.Context, payload model.UserPayload) (wallet.WalletResponse, error) {
	const layout = "2006-01-02"

	parseDate, err := time.Parse(layout, payload.DoB)
	if err != nil {
		return wallet.WalletResponse{}, fmt.Errorf("failed to parse data :%v", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return wallet.WalletResponse{}, err
	}

	user := sqlc.InsertUserParams{
		Username:    payload.Username,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Address:     payload.Address,
		Dob: pgtype.Date{
			Time:  parseDate,
			Valid: true,
		},
		Fullname: payload.Fullname,
		Password: string(password),
	}

	id, err := s.q.InsertUser(ctx, user)
	if err != nil {
		return wallet.WalletResponse{}, err
	}

	respWallet, err := s.wallet.CreateWallet(ctx, id)
	if err != nil {
		if err := s.q.DeleteUserByID(ctx, id); err != nil {
			return wallet.WalletResponse{}, err
		}
		return wallet.WalletResponse{}, err
	}

	if err := s.notif.SendNotification(ctx, external.NotifRequest{
		Recipient:    payload.Email,
		TemplateName: "register",
		Placeholder: map[string]string{
			"fullname": payload.Fullname,
		},
	}); err != nil {
		return wallet.WalletResponse{}, err
	}

	return *respWallet, nil
}

func (s *UserService) GetUser(ctx context.Context, payload model.UserLoginPayload) (*model.LoginResponse, error) {
	user, err := s.q.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, err
	}

	activeToken, err := s.auth.GenerateToken(user.ID, "active_token")
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.auth.GenerateToken(user.ID, "refresh_token")
	if err != nil {
		return nil, err
	}

	insertToken := sqlc.InsertTokenParams{
		UserID:       user.ID,
		Token:        activeToken,
		RefreshToken: refreshToken,
		TokenExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(auth.TokenTime["active_token"]),
			Valid: true,
		},
		RefreshTokenExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(auth.TokenTime["refresh_token"]),
			Valid: true,
		},
	}

	if err := s.q.InsertToken(ctx, insertToken); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		ActiveToken:  activeToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (*sqlc.User, error) {
	user, err := s.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &sqlc.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Dob:         user.Dob,
		Fullname:    user.Fullname,
	}, nil
}

func (s *UserService) DeleteTokenByID(ctx context.Context, id int32) error {
	if err := s.q.DeleteTokenByUserID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *UserService) RefreshToken(ctx context.Context, user *sqlc.User) (*model.LoginResponse, error) {
	token, err := s.q.GetTokenByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	refreshTokenExpired := token.RefreshTokenExpiresAt.Time.Truncate(time.Second)
	now := time.Now().Truncate(time.Second).UTC().Add(time.Hour * 7)

	if now.After(refreshTokenExpired) {
		if err := s.q.DeleteTokenByToken(ctx, token.Token); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("token has expired, please login again")
	}

	newToken, err := s.auth.GenerateToken(user.ID, "active_token")
	if err != nil {
		return nil, err
	}

	updateToken := sqlc.UpdateTokenParams{
		Token: newToken,
		TokenExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(auth.TokenTime["active_token"]),
			Valid: true,
		},
		UserID:       user.ID,
		RefreshToken: token.RefreshToken,
	}

	if err := s.q.UpdateToken(ctx, updateToken); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		ActiveToken:  newToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
