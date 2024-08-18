-- name: CreateChatGroup :one
INSERT INTO chat_groups (chat_id, user_id)
SELECT $1, $2
WHERE NOT EXISTS( 
    SELECT * FROM chat_groups
    WHERE chat_id = $1 AND user_id = $2
)
RETURNING *;

-- name: GetAllFromGroup :many
SELECT * FROM chat_groups
WHERE chat_id = $1;

-- name: GetAllGroupsForUser :many
SELECT * FROM chat_groups
WHERE user_id = $1;

-- name: GetAllChatNamesForUser :many
SELECT name
FROM chats
JOIN chat_groups ON chats.chat_id = chat_groups.chat_id
WHERE chat_groups.user_id = $1;

-- name: RemoveFromGroup :exec
DELETE FROM chat_groups
WHERE chat_id = $1 AND user_id = $2;

-- name: DeleteGroup :exec
DELETE FROM chat_groups
WHERE chat_id = $1;