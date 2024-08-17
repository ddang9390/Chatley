-- name: CreateUser :one
INSERT INTO users (id, Email, Password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetOneUser :one
SELECT * FROM users
WHERE $1 = id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE $1 = email AND $2 = Password;