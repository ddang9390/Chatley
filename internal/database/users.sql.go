// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, Email, Password)
VALUES ($1, $2, $3)
RETURNING id, email, password
`

type CreateUserParams struct {
	ID       string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Email, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE $1 = email AND $2 = Password
`

type DeleteUserParams struct {
	Email    string
	Password string
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser, arg.Email, arg.Password)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, password FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Email, &i.Password); err != nil {
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

const getOneUser = `-- name: GetOneUser :one
SELECT id, email, password FROM users
WHERE $1 = id
`

func (q *Queries) GetOneUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getOneUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return i, err
}
