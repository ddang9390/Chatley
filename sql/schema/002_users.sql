-- +goose Up
ALTER TABLE users
ADD UNIQUE (Email);