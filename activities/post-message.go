package activities

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/Julien4218/temporal-slack-activity/instrumentation"
	"github.com/Julien4218/temporal-slack-activity/models"
)

func PostMessageActivity(ctx context.Context, slackActivityData models.SlackActivityData) (string, error) {
	return NewSlackActivityContext().postMessageActivityImpl(ctx, slackActivityData)
}

func (c *SlackActivityContext) postMessageActivityImpl(ctx context.Context, slackActivityData models.SlackActivityData) (string, error) {

	slackAttachment := translateToSlackAttachment(slackActivityData)
	channelID, timestamp, err := c.client.PostMessage(slackActivityData.ChannelId, slack.MsgOptionText(slackActivityData.FirstResponseWarning, false), slack.MsgOptionAttachments(slackAttachment))
	if err != nil {
		instrumentation.Log(fmt.Sprintf("Error:%s", err.Error()))
		return "", err
	}

	instrumentation.Log(fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))
	return "", nil
}

func translateToSlackAttachment(slackActivityData models.SlackActivityData) slack.Attachment {
	return slack.Attachment{
		Pretext: slackActivityData.Attachment.Pretext,
		Text:    slackActivityData.Attachment.Text,
	}
}
