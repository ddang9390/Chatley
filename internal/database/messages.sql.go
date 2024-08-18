// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.sql

package database

import (
	"context"
	"database/sql"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (created_date, content, sender, chat_id)
VALUES (CURRENT_TIMESTAMP, $1, $2, $3)
RETURNING message_id, created_date, content, sender, chat_id
`

type CreateMessageParams struct {
	Content sql.NullString
	Sender  sql.NullString
	ChatID  sql.NullInt32
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessage, arg.Content, arg.Sender, arg.ChatID)
	var i Message
	err := row.Scan(
		&i.MessageID,
		&i.CreatedDate,
		&i.Content,
		&i.Sender,
		&i.ChatID,
	)
	return i, err
}

const deleteAllMessagesForUser = `-- name: DeleteAllMessagesForUser :exec
DELETE FROM messages
WHERE chat_id = $1 AND sender = $2
`

type DeleteAllMessagesForUserParams struct {
	ChatID sql.NullInt32
	Sender sql.NullString
}

func (q *Queries) DeleteAllMessagesForUser(ctx context.Context, arg DeleteAllMessagesForUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteAllMessagesForUser, arg.ChatID, arg.Sender)
	return err
}

const deleteMessageForUser = `-- name: DeleteMessageForUser :exec
DELETE FROM messages
WHERE chat_id = $1 AND sender = $2 AND message_id = $3
`

type DeleteMessageForUserParams struct {
	ChatID    sql.NullInt32
	Sender    sql.NullString
	MessageID int32
}

func (q *Queries) DeleteMessageForUser(ctx context.Context, arg DeleteMessageForUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteMessageForUser, arg.ChatID, arg.Sender, arg.MessageID)
	return err
}

const getMessagesFromChat = `-- name: GetMessagesFromChat :many
SELECT message_id, created_date, content, sender, chat_id FROM messages
WHERE chat_id = $1
ORDER BY created_date
`

func (q *Queries) GetMessagesFromChat(ctx context.Context, chatID sql.NullInt32) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessagesFromChat, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.MessageID,
			&i.CreatedDate,
			&i.Content,
			&i.Sender,
			&i.ChatID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessagesFromChatAndUser = `-- name: GetMessagesFromChatAndUser :many
SELECT message_id, created_date, content, sender, chat_id FROM messages
WHERE chat_id = $1 AND sender = $2
ORDER BY created_date
`

type GetMessagesFromChatAndUserParams struct {
	ChatID sql.NullInt32
	Sender sql.NullString
}

func (q *Queries) GetMessagesFromChatAndUser(ctx context.Context, arg GetMessagesFromChatAndUserParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessagesFromChatAndUser, arg.ChatID, arg.Sender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.MessageID,
			&i.CreatedDate,
			&i.Content,
			&i.Sender,
			&i.ChatID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
