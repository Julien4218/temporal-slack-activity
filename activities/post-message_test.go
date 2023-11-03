package activities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostMessageShouldSucceed(t *testing.T) {
	client := NewMockSlackClient()
	ctx := NewSlackActivityContextWith(client)
	result, err := ctx.postMessageActivityImpl(context.Background(), "a name", "a channel")
	require.NoError(t, err)
	require.Equal(t, result, "")
	require.Equal(t, client.PostMessageCount, 1)
}
