package envprovider

import (
	"os"

	"github.com/gosom/secretstash"
)

var _ secretstash.SecretProvider = (*envProvider)(nil)

type envProvider struct{}

func New() secretstash.SecretProvider {
	return &envProvider{}
}

func (p *envProvider) GetSecret(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", secretstash.ErrSecretNotFound
	}

	return value, nil
}
