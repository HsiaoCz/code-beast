-- +goose Up
CREATE TABLE IF NOT EXISTS messages();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS messages;
-- +goose StatementBegin
-- +goose StatementEnd
