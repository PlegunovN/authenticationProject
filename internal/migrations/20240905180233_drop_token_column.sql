-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(999);
ALTER TABLE users DROP COLUMN token ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN password;
ALTER TABLE users ALTER COLUMN token TYPE VARCHAR(999);
-- +goose StatementEnd