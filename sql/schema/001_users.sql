-- +goose Up
CREATE TABLE users(
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);


-- +goose Down
DROP Table users;