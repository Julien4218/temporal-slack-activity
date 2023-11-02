package activities

import (
	"github.com/slack-go/slack"
)

type MockSlackClient struct {
	PostMessageCount int
}

func NewMockSlackClient() *MockSlackClient {
	return &MockSlackClient{}
}

func (c *MockSlackClient) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	c.PostMessageCount++
	return "", "", nil
}
