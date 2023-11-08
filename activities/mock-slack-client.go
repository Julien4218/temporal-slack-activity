package activities

import (
	"github.com/slack-go/slack"
)

type MockSlackClient struct {
	PostMessageCount int
}

func (c *MockSlackClient) AddReaction(name string, item slack.ItemRef) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockSlackClient) GetReactions(item slack.ItemRef, params slack.GetReactionsParameters) ([]slack.ItemReaction, error) {
	//TODO implement me
	panic("implement me")
}

func NewMockSlackClient() *MockSlackClient {
	return &MockSlackClient{}
}

func (c *MockSlackClient) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	c.PostMessageCount++
	return "", "", nil
}
