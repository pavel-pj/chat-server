package chat

import (
	"context"

	"github.com/pavel-pj/go-chat/internal/model"
)

func (s ChatService) Get(ctx context.Context, id int64) (model.Chat, error) {

	data, err := s.Repository.Get(ctx, id)
	if err != nil {
		return model.Chat{}, err
	}

	return data, nil
}
