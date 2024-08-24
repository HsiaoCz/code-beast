-- +goose Up
CREATE TABLE IF NOT EXISTS push_historys();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS push_historys;
-- +goose StatementBegin
-- +goose StatementEnd
