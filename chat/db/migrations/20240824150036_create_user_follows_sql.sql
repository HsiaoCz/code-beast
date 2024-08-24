-- +goose Up
CREATE TABLE IF NOT EXISTS user_follows();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS user_follows;
-- +goose StatementBegin
-- +goose StatementEnd