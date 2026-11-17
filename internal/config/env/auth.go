package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pavel-pj/go-chat/internal/config"
)

type AuthConfig struct {
	address string
}

func NewAuthConfig() (config.AuthConfig, error) {

	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
		return AuthConfig{}, nil
	}

	return AuthConfig{
		address: getEnv("AUTH_GRPC_ADDR", "50560"),
	}, nil

}

func (a AuthConfig) Address() string {
	return a.address
}
