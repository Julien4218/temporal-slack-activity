package activities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostMessageShouldSucceed(t *testing.T) {
	result, err := PostMessageActivity(context.Background(), "input")
	require.NoError(t, err)
	require.Equal(t, *result, "pass")
}
