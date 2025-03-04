// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_session.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteTokenByToken = `-- name: DeleteTokenByToken :exec
DELETE FROM user_session WHERE token = $1
`

func (q *Queries) DeleteTokenByToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteTokenByToken, token)
	return err
}

const deleteTokenByUserID = `-- name: DeleteTokenByUserID :exec
DELETE FROM user_session WHERE user_id = $1
`

func (q *Queries) DeleteTokenByUserID(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteTokenByUserID, userID)
	return err
}

const getTokenByToken = `-- name: GetTokenByToken :one
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, created_at, updated_at
FROM user_session
WHERE token = $1
`

func (q *Queries) GetTokenByToken(ctx context.Context, token string) (UserSession, error) {
	row := q.db.QueryRow(ctx, getTokenByToken, token)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.RefreshToken,
		&i.TokenExpiresAt,
		&i.RefreshTokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTokenByUserID = `-- name: GetTokenByUserID :one
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, created_at, updated_at
FROM user_session
WHERE user_id = $1
`

func (q *Queries) GetTokenByUserID(ctx context.Context, userID int32) (UserSession, error) {
	row := q.db.QueryRow(ctx, getTokenByUserID, userID)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.RefreshToken,
		&i.TokenExpiresAt,
		&i.RefreshTokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertToken = `-- name: InsertToken :exec
INSERT INTO user_session (user_id, token, refresh_token, token_expires_at, refresh_token_expires_at) VALUES ($1, $2, $3, $4, $5)
`

type InsertTokenParams struct {
	UserID                int32
	Token                 string
	RefreshToken          string
	TokenExpiresAt        pgtype.Timestamp
	RefreshTokenExpiresAt pgtype.Timestamp
}

func (q *Queries) InsertToken(ctx context.Context, arg InsertTokenParams) error {
	_, err := q.db.Exec(ctx, insertToken,
		arg.UserID,
		arg.Token,
		arg.RefreshToken,
		arg.TokenExpiresAt,
		arg.RefreshTokenExpiresAt,
	)
	return err
}

const updateToken = `-- name: UpdateToken :exec
UPDATE user_session SET token = $1, token_expires_at = $2, updated_at = now() WHERE user_id = $3 AND refresh_token = $4
`

type UpdateTokenParams struct {
	Token          string
	TokenExpiresAt pgtype.Timestamp
	UserID         int32
	RefreshToken   string
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) error {
	_, err := q.db.Exec(ctx, updateToken,
		arg.Token,
		arg.TokenExpiresAt,
		arg.UserID,
		arg.RefreshToken,
	)
	return err
}
