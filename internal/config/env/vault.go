package env

import (
	"log"

	"github.com/joho/godotenv"
)

type VaultConfig struct {
	address  string
	token    string
	roleID   string
	secretID string

	pkiMount      string
	pkiRole       string
	pkiCommonName string
	pkiTTL        string
	rootCAFile    string
}

func NewVaultConfig() (VaultConfig, error) {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	cfg := VaultConfig{}

	cfg.address = getEnv("VAULT_ADDR", "http://127.0.0.1:8200")
	cfg.token = getEnv("VAULT_TOKEN", "")
	cfg.roleID = getEnv("VAULT_ROLE_ID", "")
	cfg.secretID = getEnv("VAULT_SECRET_ID", "")

	cfg.pkiMount = getEnv("VAULT_PKI_MOUNT", "pki")
	cfg.pkiRole = getEnv("VAULT_PKI_ROLE", "chat")
	cfg.pkiCommonName = getEnv("VAULT_PKI_COMMON_NAME", "localhost")
	cfg.pkiTTL = getEnv("VAULT_PKI_TTL", "720h")
	cfg.rootCAFile = getEnv("VAULT_ROOT_CA_FILE", "")

	return cfg, nil
}

func (s VaultConfig) Address() string {
	return s.address
}

func (s VaultConfig) Token() string {
	return s.token
}

func (s VaultConfig) RoleID() string {
	return s.roleID
}

func (s VaultConfig) SecretID() string {
	return s.secretID
}

func (s VaultConfig) PKIMount() string {
	return s.pkiMount
}

func (s VaultConfig) PKIRole() string {
	return s.pkiRole
}

func (s VaultConfig) PKICommonName() string {
	return s.pkiCommonName
}

func (s VaultConfig) PKITTL() string {
	return s.pkiTTL
}

func (s VaultConfig) RootCAFile() string {
	return s.rootCAFile
}
