-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT loginUnic UNIQUE (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
