// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Chat struct {
	ChatID      int32
	CreatedDate time.Time
	Name        string
}

type ChatGroup struct {
	ChatGroupID int32
	ChatID      int32
	UserID      string
}

type Message struct {
	MessageID   int32
	CreatedDate time.Time
	Content     sql.NullString
	Sender      sql.NullString
	ChatID      sql.NullInt32
}

type User struct {
	ID          string
	Email       string
	Password    string
	CreatedDate time.Time
}
