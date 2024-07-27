-- +goose Up
-- +goose StatementBegin
DROP TABLE usersandtokens;
DROP TABLE tokens;
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(999);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE usersandtokens;
DROP TABLE tokens;
DROP TABLE users;
-- +goose StatementEnd
