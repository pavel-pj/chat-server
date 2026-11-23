package chat

import (
	"context"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
)

func (s ChatService) Get(ctx context.Context, id int64) (*desc.Chat, error) {

	data, err := s.Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
