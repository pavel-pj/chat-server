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

	PKIMount() string
	PKIRole() string
	PKICommonName() string
	PKITTL() string
	RootCAFile() string
}

type AuthConfig interface {
	Address() string
	TLSServerName() string
}

type LoggerConfig interface {
	Level() string
	Filename() string
	MaxSize() int
	MaxBackups() int
	MaxAge() int
}
