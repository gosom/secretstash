package secretstash

import (
	"errors"
	"fmt"
)

var (
	ErrSecretNotFound = errors.New("secret not found")
	ErrAtLeastOne     = errors.New("at least one secret provider must be provided")
)

// SecretProvider is an interface that can be implemented by any type that can
// provide a secret.
//
//go:generate mockgen -destination=./mocks/mock_secret_provider.go -package=mocks . SecretProvider
type SecretProvider interface {
	GetSecret(name string) (string, error)
}

// SecretStash is a collection of SecretProviders that can be queried for a
// secret.
type SecretStash struct {
	providers []SecretProvider
}

// New returns a new SecretStash with the given providers.
func New(providers ...SecretProvider) (*SecretStash, error) {
	if len(providers) == 0 {
		return nil, ErrAtLeastOne
	}

	ans := SecretStash{
		providers: providers,
	}

	return &ans, nil
}

// GetSecret returns the secret with the given name from the first provider that
// can provide it. If no provider can provide the secret, an ErrSecretNotFound error is returned.
func (s *SecretStash) GetSecret(name string) (string, error) {
	for _, provider := range s.providers {
		secret, err := provider.GetSecret(name)
		if err == nil {
			return secret, nil
		}

		if !errors.Is(err, ErrSecretNotFound) {
			return "", fmt.Errorf("%s: %w", name, err)
		}
	}

	return "", fmt.Errorf("%s: %w", name, ErrSecretNotFound)
}
