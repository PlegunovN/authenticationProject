-- +goose Up
-- +goose StatementBegin
DROP TABLE usersandtokens;
DROP TABLE tokens;
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(999);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE usersandtokens(
id BIGSERIAL PRIMARY KEY,
userID BIGINT REFERENCES users,
tokenID BIGINT REFERENCES tokens
);
CREATE TABLE tokens(
id BIGSERIAL PRIMARY KEY,
token VARCHAR(256)
);
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(25);
-- +goose StatementEnd
