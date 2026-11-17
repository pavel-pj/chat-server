package app

import (
	//"github.com/pavel-pj/auth/internal/closer"
	"context"
	"log"

	"github.com/pavel-pj/go-chat/internal/config"

	env "github.com/pavel-pj/go-chat/internal/config/env"

	repository "github.com/pavel-pj/go-chat/internal/repository"
	chatRepo "github.com/pavel-pj/go-chat/internal/repository/chat"

	"github.com/pavel-pj/go-chat/internal/service"
	chatService "github.com/pavel-pj/go-chat/internal/service/chat"

	db "github.com/pavel-pj/platform_lib/pkg/client/db"
	"github.com/pavel-pj/platform_lib/pkg/client/db/pg"

	//pg "github.com/pavel-pj/platform_lib/pkg/client/db/pg"
	"github.com/pavel-pj/platform_lib/pkg/closer"

	grpcChatApi "github.com/pavel-pj/go-chat/internal/api/chat"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	authConfig  config.AuthConfig
	grpcConfig  config.GRPCConfig
	vaultConfig config.VaultConfig

	//pgPool *pgxpool.Pool
	pgClient db.Client
	//txManager db.TxManager

	chatImpl       *grpcChatApi.ChatImplementation
	chatService    service.ChatService
	chatRepository repository.ChatRepository
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {

	if s.pgConfig == nil {
		conf, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = conf
	}

	return s.pgConfig

}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {

	if s.grpcConfig == nil {
		conf, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = conf
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetAuthConfig() config.AuthConfig {

	if s.authConfig == nil {
		conf, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.authConfig = conf
	}

	return s.authConfig
}

func (s *serviceProvider) GetVaultConfig() config.VaultConfig {
	if s.vaultConfig == nil {
		conf, err := env.NewVaultConfig()
		if err != nil {
			log.Fatalf("failed to get vault config: %s", err.Error())
		}
		s.vaultConfig = conf
	}

	return s.vaultConfig
}

func (s *serviceProvider) GetDBClient(ctx context.Context) db.Client {
	if s.pgClient == nil {

		client, err := pg.New(context.Background(), s.GetPGConfig().DNS())

		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}

		// Проверяем подключение
		err = client.DB().Ping(context.Background())
		if err != nil {
			log.Fatalf("failed to ping database: %s", err.Error())
		}

		closer.Add(func() error {
			// тк pool.Close не возвращает error - создаем обертку
			client.Close()
			return nil
		})

		s.pgClient = client

	}

	return s.pgClient
}

func (s *serviceProvider) GetChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewChatRepository(s.GetDBClient(ctx))

	}
	return s.chatRepository
}

func (s *serviceProvider) GetChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(s.GetChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) GetChatImplimentation(ctx context.Context) *grpcChatApi.ChatImplementation {
	if s.chatImpl == nil {
		s.chatImpl = grpcChatApi.NewChatImplimentation(s.GetChatService(ctx))
	}

	return s.chatImpl
}
