-- +goose Up
CREATE TABLE IF NOT EXISTS sessions(
    id integer primary key,
    token text not null,
    user_id text not null references users,
    ip_address text,
    user_agent text,
    expires_at datetime not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS sessions;
-- +goose StatementBegin
-- +goose StatementEnd