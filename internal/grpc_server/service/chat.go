package service

import (
	repo "chat-server/internal/grpc_server/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Repository repo.Repository
}

func NewService(DB *pgxpool.Pool) Service {

	service := repo.NewRepository(DB)

	return Service{
		Repository: service,
	}
}

func (s *Service) Store(ctx context.Context, users []string) (int64, error) {

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
