package service

import (
	repository "chat-server/internal/grpc_server/repository"
	chatRepository "chat-server/internal/grpc_server/repository/chat"
	"context"

	desc "chat-server/pkg/chat_v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Repository repository.ChatRepository
}

func NewService(DB *pgxpool.Pool) Service {

	service := chatRepository.NewRepository(DB)

	return Service{
		Repository: service,
	}
}

func (s *Service) Store(ctx context.Context, users *desc.ChatInfo) (int64, error) {

	result, err := s.Repository.Store(ctx, users)
	if err != nil {
		return 0, err
	}
	return result, err
}

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
