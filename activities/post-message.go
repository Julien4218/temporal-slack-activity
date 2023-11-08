package activities

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/Julien4218/temporal-slack-activity/instrumentation"
	"github.com/Julien4218/temporal-slack-activity/models"
)

func PostMessageActivity(ctx context.Context, slackActivityData models.SlackActivityData) (models.MessageDetails, error) {
	return NewSlackActivityContext().postMessageActivityImpl(ctx, slackActivityData)
}

func (c *SlackActivityContext) postMessageActivityImpl(ctx context.Context, slackActivityData models.SlackActivityData) (models.MessageDetails, error) {

	slackAttachment := translateToSlackAttachment(slackActivityData)
	responseChannelID, responseTimestamp, err := c.client.PostMessage(slackActivityData.ChannelId, slack.MsgOptionText(slackActivityData.FirstResponseWarning, false), slack.MsgOptionAttachments(slackAttachment))

	instrumentation.Log(fmt.Sprintf("Message successfully sent to channel %s at %s", responseChannelID, responseTimestamp))
	return models.MessageDetails{
		ChannelID: responseChannelID,
		Timestamp: responseTimestamp,
	}, err
}

func translateToSlackAttachment(slackActivityData models.SlackActivityData) slack.Attachment {
	return slack.Attachment{
		Pretext: slackActivityData.Attachment.Pretext,
		Text:    slackActivityData.Attachment.Text,
	}
}
