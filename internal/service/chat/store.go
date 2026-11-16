package chat

import (
	"context"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
)

func (s ChatService) Store(ctx context.Context, users *desc.ChatInfo) (int64, error) {

	result, err := s.Repository.Store(ctx, users)
	if err != nil {
		return 0, err
	}
	return result, err
}
