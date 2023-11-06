package activities

import (
	"context"
	"testing"

	"github.com/Julien4218/temporal-slack-activity/models"

	"github.com/stretchr/testify/require"
)

func TestPostMessageShouldSucceed(t *testing.T) {
	client := NewMockSlackClient()
	ctx := NewSlackActivityContextWith(client)
	attachment := models.MessageAttachment{
		Pretext: "foo",
		Text:    "bar",
	}
	result, err := ctx.postMessageActivityImpl(context.Background(), "a name", "a channel", attachment)
	require.NoError(t, err)
	require.Equal(t, result, "")
	require.Equal(t, client.PostMessageCount, 1)
}
