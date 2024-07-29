-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN token VARCHAR(999);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
