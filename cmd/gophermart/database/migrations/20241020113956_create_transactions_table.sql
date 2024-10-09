-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
     id BIGSERIAL PRIMARY KEY,
     amount BIGINT NOT NULL,
     user_id BIGINT NOT NULL,
     order_id BIGINT NULL,
     processed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     FOREIGN KEY(user_id) 
       REFERENCES users(id)
       ON DELETE SET NULL,
    FOREIGN KEY(order_id) 
       REFERENCES orders(id)
       ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
