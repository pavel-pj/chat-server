-- +goose Up
-- +goose StatementBegin
CREATE TABLE chat_users (
    id BIGSERIAL PRIMARY KEY,
    user_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat_users;
-- +goose StatementEnd
