package activities

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/Julien4218/temporal-slack-activity/instrumentation"
	"github.com/Julien4218/temporal-slack-activity/models"
)

func PostMessageActivity(ctx context.Context, firstResponseWarning string, channel string, attachment models.MessageAttachment) (string, error) {
	return NewSlackActivityContext().postMessageActivityImpl(ctx, firstResponseWarning, channel, attachment)
}

func (c *SlackActivityContext) postMessageActivityImpl(ctx context.Context, firstResponseWarning string, channel string, attachment models.MessageAttachment) (string, error) {

	slackAttachment := slack.Attachment{
		Pretext: attachment.Pretext,
		Text:    attachment.Text,
	}
	instrumentation.Log(fmt.Sprintf("SlackMessageActivity firstResponseWarning:%s", firstResponseWarning))
	channelID, timestamp, err := c.client.PostMessage(channel, slack.MsgOptionText(firstResponseWarning, false), slack.MsgOptionAttachments(slackAttachment))
	if err != nil {
		instrumentation.Log(fmt.Sprintf("Error:%s", err.Error()))
		return "", err
	}

	instrumentation.Log(fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))
	return "", nil
}
