package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pavel-pj/go-chat/internal/config/env"
	"github.com/pavel-pj/go-chat/internal/vault"
	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	"github.com/pavel-pj/platform_lib/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// App хранит ВСЕ зависимости приложения
type App struct {
	//Services serv.UserService
	//Pool     *pgxpool.Pool
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
	vault           *vault.Secret
}

func NewApp(ctx context.Context) (*App, error) {
	// Создаем пул соединений pgx

	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := a.runGRPCServer()
	if err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}

	return nil

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initVault,
		a.initGRPCServer,
		//a.initHTTPServer,
		//a.initSwaggerServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := env.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = NewServiceProvider()
	return nil
}

func (a *App) initVault(ctx context.Context) error {
	cfg := a.serviceProvider.GetVaultConfig()

	certs, err := vault.Load(ctx, vault.Config{
		Address:    cfg.Address(),
		Token:      cfg.Token(),
		RoleID:     cfg.RoleID(),
		SecretID:   cfg.SecretID(),
		Mount:      cfg.KVMount(),
		SecretPath: cfg.SecretPath(),
		ServerName: cfg.TLSServerName(),
	})
	if err != nil {
		return fmt.Errorf("init vault: %w", err)
	}

	a.vault = certs

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {

	a.grpcServer = grpc.NewServer(
		grpc.Creds(credentials.NewTLS(a.vault.TLS.ServerTLS)),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.ChainUnaryInterceptor(
		//	interceptor.ValidateInterceptor,
		//	interceptor.RecoveryInterceptor),
	)
	// Только для тестовой среды - для удобства работы в Postman
	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.GetChatImplimentation(ctx))
	//descAuth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.GetAuthImplimentation(ctx))
	//descAccess.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.GetAccessImplimentation(ctx))

	closer.Add(func() error {
		a.grpcServer.GracefulStop()
		return nil
	})

	return nil
}

func (a *App) runGRPCServer() error {

	list, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().Address())
	if err != nil {
		return err
	}
	log.Printf("GRPC Server is running in %s", a.serviceProvider.GetGRPCConfig().Address())
	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil

}
