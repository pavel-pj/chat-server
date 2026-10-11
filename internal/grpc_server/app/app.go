package app

import (
	"context"
	//"chat-server/internal/db/generated"
	//"chat-server/internal/services"
	config "chat-server/internal/grpc_server/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// App хранит ВСЕ зависимости приложения
type App struct {
	DB *pgxpool.Pool
	//Queries  *generated.Queries
	//Services *Service
}

//type Service struct {
//	Links *services.LinkService
//}

func NewApp(cfg config.Config) (*App, error) {
	// Создаем пул соединений pgx

	dbURL := "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable"

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("SUPER ERROR failed to ping database: %w", err)
	}

	// Создаем queries (теперь совместимо!)
	//queries := generated.New(pool)

	//services := &Service{
	//	Links: services.NewLinkService(queries),
	//}

	return &App{
		DB: pool,
		//Queries:  queries,
		//Services: services,
	}, nil
}

// Close закрывает соединение с БД
func (a *App) Close() error {
	if a.DB != nil {
		a.DB.Close()
	}
	return nil
}
