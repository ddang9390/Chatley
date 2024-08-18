-- +goose Up
CREATE TABLE chats(
    chat_id SERIAL NOT NULL PRIMARY KEY,
    created_date TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL
);


-- +goose Down
DROP Table chats;