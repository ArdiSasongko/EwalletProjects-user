package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	q *sqlc.Queries
}

func (s *UserService) InsertUser(ctx context.Context, payload model.UserPayload) error {
	const layout = "2006-01-02"

	parseDate, err := time.Parse(layout, payload.DoB)
	if err != nil {
		return fmt.Errorf("failed to parse data :%v", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
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

	_, err = s.q.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(ctx context.Context) error {
	return nil
}

func (s *UserService) insertToken(ctx context.Context) error {
	return nil
}
