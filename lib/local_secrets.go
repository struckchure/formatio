package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type LocalSecrets struct{}

// Load implements SecretsInterface.
func (LocalSecrets) Load() error {
	return godotenv.Load()
}

// Get implements SecretsInterface.
func (LocalSecrets) Get(key string) string {
	return os.Getenv(key)
}

// GetOrPanic implements SecretsInterface.
func (l LocalSecrets) GetOrPanic(key string) string {
	value := l.Get(key)
	if value == "" {
		log.Panicf("%s is empty", key)
	}

	return value
}

func (*LocalSecrets) NewLocalSecret() ISecrets {
	return LocalSecrets{}
}
