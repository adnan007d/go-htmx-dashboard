-- name: CreateUser :one
INSERT INTO users (id, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
