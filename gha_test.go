package ghasign

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/openpubkey/openpubkey/client"
	"github.com/openpubkey/openpubkey/providers"
	"github.com/openpubkey/openpubkey/verifier"
	"github.com/stretchr/testify/require"
)

func TestGithub(t *testing.T) {
	var op client.OpenIdProvider
	var err error

	op, err = providers.NewGithubOpFromEnvironment()
	require.NoError(t, err)

	c, err := client.New(op)
	require.NoError(t, err)

	pkt, err := c.Auth(context.TODO())
	require.NoError(t, err)
	require.NotNil(t, pkt)
	fmt.Println("New PK token generated")

	// For debugging purposes, lets us write the PK Token to the test log
	outputPktToken := false
	if outputPktToken {
		pktCom, _ := pkt.Compact()
		b64pktCom := base64.StdEncoding.EncodeToString(pktCom)
		fmt.Println(string(b64pktCom))
		require.False(t, true) // Println in tests are hidden unless test fails.
	}

	ver, err := verifier.New(op)
	require.NoError(t, err)

	err = ver.VerifyPKToken(context.TODO(), pkt)
	require.NoError(t, err)

}
