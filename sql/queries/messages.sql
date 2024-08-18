-- name: CreateMessage :one
INSERT INTO messages (created_date, content, sender, chat_id)
VALUES (CURRENT_TIMESTAMP, $1, $2, $3)
RETURNING *;


-- name: GetMessagesFromChat :many
SELECT * FROM messages
WHERE chat_id = $1
ORDER BY created_date
LIMIT 100;

-- name: GetMessagesFromChatAndUser :many
SELECT * FROM messages
WHERE chat_id = $1 AND sender = $2
ORDER BY created_date
LIMIT 100;

-- name: DeleteMessageForUser :exec
DELETE FROM messages
WHERE chat_id = $1 AND sender = $2 AND message_id = $3;

-- name: DeleteAllMessagesForUser :exec
DELETE FROM messages
WHERE chat_id = $1 AND sender = $2;

-- name: EditMessage :exec
UPDATE messages
SET content = $1
WHERE chat_id = $2 AND sender = $3;