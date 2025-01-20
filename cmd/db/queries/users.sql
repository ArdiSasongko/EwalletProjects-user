-- name: InsertUser :one
INSERT INTO users (username, email, phone_number, address, dob, password, fullname)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetUserByID :one
SELECT id, username, email, phone_number, address, dob, fullname, password, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT id, username, email, phone_number, address, dob, fullname, password, created_at, updated_at
FROM users
WHERE username = $1;