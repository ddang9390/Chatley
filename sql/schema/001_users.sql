-- +goose Up
CREATE TABLE Users(
    Id VARCHAR(255) NOT NULL PRIMARY KEY,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL
);


-- +goose Down
DROP Table Users;