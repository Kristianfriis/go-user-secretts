package config

import "github.com/google/uuid"

type Config struct {
	LocalSecretFileName string
	UserSecretFileName  string
	UserSecretsFolder   string
	UserSecretId        uuid.UUID
}

func NewConfig() Config {
	return Config{
		LocalSecretFileName: "user-secret.config",
		UserSecretFileName:  "user-secret.env",
		UserSecretsFolder:   "go-user-secrets",
		UserSecretId:        uuid.New(),
	}
}
