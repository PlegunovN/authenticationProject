-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY ,
    login VARCHAR(25),
    password VARCHAR(25)
);

CREATE TABLE tokens (
    id BIGSERIAL PRIMARY KEY,
    token VARCHAR(256)
);

CREATE TABLE usersandtokens (
    id BIGSERIAL PRIMARY KEY,
    userID BIGINT REFERENCES users,
    tokenID BIGINT REFERENCES tokens
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE tokens;
DROP TABLE usersandtokens;
-- +goose StatementEnd
