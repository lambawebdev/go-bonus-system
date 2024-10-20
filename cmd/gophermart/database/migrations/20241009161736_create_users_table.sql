-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
     id BIGSERIAL PRIMARY KEY,
     login VARCHAR(255) UNIQUE,
     password VARCHAR(255)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
