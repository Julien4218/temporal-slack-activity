package activities

import "github.com/slack-go/slack"

type SlackClient interface {
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
}
