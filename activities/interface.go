package activities

import "github.com/slack-go/slack"

type SlackClient interface {
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
	GetReactions(item slack.ItemRef, params slack.GetReactionsParameters) ([]slack.ItemReaction, error)
}
