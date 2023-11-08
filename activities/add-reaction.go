package activities

import (
	"context"
	"github.com/slack-go/slack"
)

func AddReaction(ctx context.Context, name string, channelID string, messageTimestamp string) error {
	return NewSlackActivityContext().addReactionImpl(ctx, name, channelID, messageTimestamp)
}

func (c *SlackActivityContext) addReactionImpl(ctx context.Context, name string, channelID string, messageTimestamp string) error {
	item := slack.ItemRef{
		Channel:   channelID,
		Timestamp: messageTimestamp,
	}
	err := c.client.AddReaction(name, item)
	return err
}
