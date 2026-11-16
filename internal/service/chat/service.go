package chat

import (
	//"context"

	"github.com/pavel-pj/go-chat/internal/repository"
	"github.com/pavel-pj/go-chat/internal/service"
	//desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
)

type ChatService struct {
	Repository repository.ChatRepository
}

func NewChatService(repository repository.ChatRepository) service.ChatService {

	return ChatService{
		Repository: repository,
	}
}

/*
func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.Repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Get(ctx context.Context, id int64) (*desc.Chat, error) {

	data, err := s.Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
*/
