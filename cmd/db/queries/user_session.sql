-- name: InsertToken :exec
INSERT INTO user_session (user_id, token, refresh_token, token_expires_at, refresh_token_expires_at) VALUES ($1, $2, $3, $4, $5);

-- name: GetTokenByUserID :one
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, created_at, updated_at
FROM user_session
WHERE user_id = $1;

-- name: GetTokenByToken :one
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, created_at, updated_at
FROM user_session
WHERE token = $1;

-- name: DeleteTokenByToken :exec
DELETE FROM user_session WHERE token = $1;

-- name: DeleteTokenByUserID :exec
DELETE FROM user_session WHERE user_id = $1;