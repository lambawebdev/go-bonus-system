-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS withdrawals (
     id BIGSERIAL PRIMARY KEY,
     sum BIGINT NOT NULL,
     user_id BIGINT NOT NULL,
     number VARCHAR(255) NOT NULL UNIQUE,
     processed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(user_id) 
       REFERENCES users(id)
       ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE withdrawals;
-- +goose StatementEnd
