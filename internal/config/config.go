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
