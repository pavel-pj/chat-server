-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    chat_users_id BIGINT NOT NULL REFERENCES chat_users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd
