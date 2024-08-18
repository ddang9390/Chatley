-- name: CreateChat :one
INSERT INTO chats (name, created_date)
VALUES ($1, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeleteChat :exec
DELETE FROM chats
WHERE chat_id = $1;