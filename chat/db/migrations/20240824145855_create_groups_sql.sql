-- +goose Up
CREATE TABLE IF NOT EXISTS groups();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS groups;
-- +goose StatementBegin
-- +goose StatementEnd
