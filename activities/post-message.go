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
	responseChannelID, responseTimestamp, err := c.client.PostMessage(slackActivityData.ChannelId,
		slack.MsgOptionText(slackActivityData.FirstResponseWarning, false),
		slack.MsgOptionAttachments(slackAttachment),
		slack.MsgOptionIconURL("https://lh4.googleusercontent.com/qyrA5kR-lTStynpecJwuLKGjrGGV1_6HJ4By4hOYLuVSq4jsm5Q3jgXr-WQHExldtD9WYO4PoEz23IqAeFaExH2UnCDphnutaI7mv3va5VPaRUyS9V-S0NaKNocj6JsIsO6km9E"),
	)
	if err != nil {
		message := fmt.Sprintf("error while posting message, detail:%s", err.Error())
		instrumentation.Log(message)
		return models.MessageDetails{}, err
	}

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
