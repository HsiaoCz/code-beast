-- +goose Up
CREATE TABLE IF NOT EXISTS applys();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS applys;
-- +goose StatementBegin
-- +goose StatementEnd
