package lib

import (
	"fmt"
)

type ISecrets interface {
	Load() error
	Get(key string) string
	GetOrPanic(key string) string
}

func SecretsFactory(factory string) (ISecrets, error) {
	var factoryObj ISecrets

	if factory == "local" {
		factoryObj = LocalSecrets{}
	} else if factory == "aws" {
		factoryObj = AwsSecrets{}
	}

	if factoryObj != nil {
		factoryObj.Load()

		return factoryObj, nil
	}

	return nil, fmt.Errorf("no secrets factory for %s", factory)
}
