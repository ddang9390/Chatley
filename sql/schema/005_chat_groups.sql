-- +goose Up
CREATE TABLE chat_groups(
    chat_group_id SERIAL NOT NULL PRIMARY KEY,
    chat_id INTEGER NOT NULL,
    user_id VARCHAR(255) NOT NULL,

    CONSTRAINT unique_chat_user UNIQUE(chat_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chats(chat_id)
);

-- +goose Down
DROP TABLE chat_groups;