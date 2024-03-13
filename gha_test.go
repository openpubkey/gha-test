package ghasign

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/openpubkey/openpubkey/client"
	"github.com/openpubkey/openpubkey/client/mocks"
	"github.com/openpubkey/openpubkey/client/providers"
	"github.com/stretchr/testify/require"
)

func TestGithub(t *testing.T) {
	var op client.OpenIdProvider
	var err error

	if RunningInGithubActions() {
		op, err = providers.NewGithubOpFromEnvironment()
		require.NoError(t, err)
	} else {
		op, err = mocks.NewMockOpenIdProvider(t, map[string]any{})
		require.NoError(t, err)
	}

	client, err := client.New(op, client.WithSignGQ(true))
	require.NoError(t, err)

	pkt, err := client.Auth(context.TODO())
	require.NoError(t, err)
	require.NotNil(t, pkt)
	fmt.Println("New PK token generated")

	err = op.Verifier().VerifyProvider(context.TODO(), pkt)
	require.NoError(t, err)
}

func RunningInGithubActions() bool {
	_, ok := os.LookupEnv("GITHUB_RUN_ID")
	return ok
}
