-- name: CreateUser :one
INSERT INTO users (id, Email, Password)
VALUES ($1, $2, $3)
RETURNING *;