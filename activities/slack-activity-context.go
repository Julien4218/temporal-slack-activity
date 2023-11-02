package activities

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackActivityContext struct {
	client SlackClient
}

func NewSlackActivityContext() *SlackActivityContext {
	return NewSlackActivityContextWith(slack.New(os.Getenv("SLACK_TOKEN")))
}

func NewSlackActivityContextWith(s SlackClient) *SlackActivityContext {
	return &SlackActivityContext{
		client: s,
	}
}
