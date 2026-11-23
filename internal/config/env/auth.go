package env

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/pavel-pj/go-chat/internal/config"
)

type AuthConfig struct {
	address       string
	tlsServerName string
}

func NewAuthConfig() (config.AuthConfig, error) {

	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
		return AuthConfig{}, nil
	}
	host := getEnv("AUTH_GRPC_HOST", "localhost")
	port := getEnv("AUTH_GRPC_PORT", "50560")

	return AuthConfig{
		address:       net.JoinHostPort(host, port),
		tlsServerName: getEnv("AUTH_TLS_SERVER_NAME", host),
	}, nil

}

func (a AuthConfig) Address() string {
	return a.address
}

func (a AuthConfig) TLSServerName() string {
	return a.tlsServerName
}
