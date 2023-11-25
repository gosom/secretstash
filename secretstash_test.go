package secretstash_test

import (
	"errors"
	"testing"

	"github.com/gosom/secretstash"
	"github.com/gosom/secretstash/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("should return an error if no providers are provided", func(t *testing.T) {
		t.Parallel()

		_, err := secretstash.New()

		require.Error(t, err)
		require.ErrorIs(t, err, secretstash.ErrAtLeastOne)
	})

	t.Run("should return a SecretStash if at least one provider is provided", func(t *testing.T) {
		mctrl := gomock.NewController(t)
		defer mctrl.Finish()

		m, err := secretstash.New(&mocks.MockSecretProvider{})
		require.NoError(t, err)

		require.NotNil(t, m)
	})
}

func Test_GetSecret(t *testing.T) {
	t.Parallel()

	t.Run("should return the secret from the first provider that can provide it", func(t *testing.T) {
		t.Parallel()

		mctrl := gomock.NewController(t)
		defer mctrl.Finish()

		p1 := mocks.NewMockSecretProvider(mctrl)
		p2 := mocks.NewMockSecretProvider(mctrl)

		stash, err := secretstash.New(p1, p2)
		require.NoError(t, err)

		p1.EXPECT().GetSecret("foo").Return("bar", nil).Times(1)

		secret, err := stash.GetSecret("foo")
		require.NoError(t, err)

		require.Equal(t, "bar", secret)
	})

	t.Run("should return the secret from the first provider that can provide it (first fails)", func(t *testing.T) {
		t.Parallel()

		mctrl := gomock.NewController(t)
		defer mctrl.Finish()

		p1 := mocks.NewMockSecretProvider(mctrl)
		p2 := mocks.NewMockSecretProvider(mctrl)

		stash, err := secretstash.New(p1, p2)
		require.NoError(t, err)

		p1.EXPECT().GetSecret("foo").Return("", secretstash.ErrSecretNotFound).Times(1)
		p2.EXPECT().GetSecret("foo").Return("bar", nil).Times(1)

		secret, err := stash.GetSecret("foo")
		require.NoError(t, err)

		require.Equal(t, "bar", secret)
	})

	t.Run("secret does not exist in any provider", func(t *testing.T) {
		t.Parallel()

		mctrl := gomock.NewController(t)
		defer mctrl.Finish()

		p1 := mocks.NewMockSecretProvider(mctrl)
		p2 := mocks.NewMockSecretProvider(mctrl)

		stash, err := secretstash.New(p1, p2)
		require.NoError(t, err)

		p1.EXPECT().GetSecret("foo").Return("", secretstash.ErrSecretNotFound).Times(1)
		p2.EXPECT().GetSecret("foo").Return("", secretstash.ErrSecretNotFound).Times(1)

		_, err = stash.GetSecret("foo")
		require.Error(t, err)
		require.ErrorIs(t, err, secretstash.ErrSecretNotFound)
	})

	t.Run("when a provider returns an error other than ErrSecretNotFound", func(t *testing.T) {
		t.Parallel()

		mctrl := gomock.NewController(t)
		defer mctrl.Finish()

		p1 := mocks.NewMockSecretProvider(mctrl)
		p2 := mocks.NewMockSecretProvider(mctrl)

		stash, err := secretstash.New(p1, p2)
		require.NoError(t, err)

		p1.EXPECT().GetSecret("foo").Return("", errors.New("any other error")).Times(1)

		_, err = stash.GetSecret("foo")
		require.Error(t, err)
		require.True(t, err.Error() == "foo: any other error")
	})
}
