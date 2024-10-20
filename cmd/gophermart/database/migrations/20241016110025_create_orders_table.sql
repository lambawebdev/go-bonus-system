-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
     id BIGSERIAL PRIMARY KEY,
     user_id BIGINT NOT NULL,
     number VARCHAR(255) NOT NULL UNIQUE,
     status INT NOT NULL DEFAULT 0,
     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(user_id) 
       REFERENCES users(id)
       ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
