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
	slackActivityData := models.SlackActivityData{
		ChannelId:            "foo",
		FirstResponseWarning: "bar",
		Attachment: models.MessageAttachment{
			Pretext: "foo",
			Text:    "bar",
		},
	}
	messageDetails, err := ctx.postMessageActivityImpl(context.Background(), slackActivityData)
	require.NoError(t, err)
	require.Equal(t, messageDetails.ChannelID, "")
	require.Equal(t, messageDetails.Timestamp, "")
	require.Equal(t, client.PostMessageCount, 1)
}
