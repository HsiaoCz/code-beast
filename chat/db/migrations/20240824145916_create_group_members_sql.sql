-- +goose Up
CREATE TABLE IF NOT EXISTS group_members();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS group_members;
-- +goose StatementBegin
-- +goose StatementEnd
