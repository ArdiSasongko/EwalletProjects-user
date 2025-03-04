// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          int32
	Username    string
	Email       string
	PhoneNumber string
	Address     string
	Dob         pgtype.Date
	Password    string
	Fullname    string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type UserSession struct {
	ID                    int32
	UserID                int32
	Token                 string
	RefreshToken          string
	TokenExpiresAt        pgtype.Timestamp
	RefreshTokenExpiresAt pgtype.Timestamp
	CreatedAt             pgtype.Timestamptz
	UpdatedAt             pgtype.Timestamptz
}
