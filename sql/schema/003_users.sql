-- +goose Up
ALTER TABLE users
ADD created_date TIMESTAMP NOT NULL;