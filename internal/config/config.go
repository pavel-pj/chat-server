package config

type PGConfig interface {
	DNS() string
}

type GRPCConfig interface {
	Address() string
}

type VaultConfig interface {
	Address() string
	Token() string
	RoleID() string
	SecretID() string
	KVMount() string
	SecretPath() string
	TLSServerName() string
}

type AuthConfig interface {
	Address() string
}
