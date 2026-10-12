package repository

import (
	desc "chat-server/pkg/chat_v1"
	"context"
)

type ChatRepository interface {
	Store(ctx context.Context, users *desc.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*desc.Chat, error)
}
