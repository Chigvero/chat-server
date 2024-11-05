-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE chats(
    id serial primary key ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE chatUsers(
    chat_id int references chats(id),
    user_name varchar(30),
    primary key (chat_id,user_name)
);

CREATE TABLE messages(
    id serial,
    chat_id int references chats(id),
    from_user_id int ,
    message_text TEXT,
    timestamp TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
drop table chatUsers;
drop table chats;
SELECT 'down SQL query';
-- +goose StatementEnd

