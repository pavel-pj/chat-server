package repository

import (
	"context"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
)

type ChatRepository interface {
	Store(ctx context.Context, users *desc.ChatInfo) (int64, error)
	//Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*desc.Chat, error)
}
