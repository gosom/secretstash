package envprovider_test

import (
	"os"
	"testing"

	"github.com/gosom/secretstash"
	"github.com/gosom/secretstash/envprovider"
	"github.com/stretchr/testify/require"
)

func Test_EnvProvider_GetSecret(t *testing.T) {
	t.Parallel()

	provider := envprovider.New()

	t.Run("should return the secret if exists in environment", func(t *testing.T) {
		t.Parallel()

		os.Setenv("foo", "bar")
		os.Setenv("foo3", "")

		secret, err := provider.GetSecret("foo")
		require.NoError(t, err)

		require.Equal(t, "bar", secret)

		secret, err = provider.GetSecret("foo3")
		require.NoError(t, err)

		require.Equal(t, "", secret)
	})

	t.Run("should return an error if the secret does not exist in environment", func(t *testing.T) {
		t.Parallel()

		_, err := provider.GetSecret("foo2")
		require.Error(t, err)
		require.ErrorIs(t, err, secretstash.ErrSecretNotFound)
	})
}
