-- name: CreateUser :one
INSERT INTO users (id, Email, Password, created_date)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetOneUser :one
SELECT * FROM users
WHERE $1 = id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE $1 = email AND $2 = Password;

-- name: UpdateUser :exec
UPDATE users
SET email = $2, Password = $3
WHERE $1 = id;