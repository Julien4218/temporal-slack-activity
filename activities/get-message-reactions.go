package activities

import (
	"context"

	"github.com/slack-go/slack"
)

func GetMessageReactions(ctx context.Context, channelID string, messageTimestamp string) ([]slack.ItemReaction, error) {
	return NewSlackActivityContext().getMessageReactionsImpl(ctx, channelID, messageTimestamp)
}

func (c *SlackActivityContext) getMessageReactionsImpl(ctx context.Context, channelID string, messageTimestamp string) ([]slack.ItemReaction, error) {
	itemRef := slack.ItemRef{
		Channel:   channelID,
		Timestamp: messageTimestamp,
	}
	// hard coding to get full reaction details. we may not need all the details and can maybe set this to false in the future or add another parameter to the method
	params := slack.GetReactionsParameters{
		Full: true,
	}
	response, err := c.client.GetReactions(itemRef, params)
	return response, err
}
