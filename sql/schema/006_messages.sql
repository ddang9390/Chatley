-- +goose Up
CREATE TABLE messages(
    message_id SERIAL NOT NULL PRIMARY KEY,
    created_date TIMESTAMP NOT NULL,
    content TEXT,

    sender VARCHAR(255) REFERENCES users(id),
    chat_id INTEGER REFERENCES chats(chat_id)
);

-- +goose Down
DROP TABLE messages;