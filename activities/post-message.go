package activities

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"

	"github.com/Julien4218/temporal-slack-activity/instrumentation"
)

func PostMessageActivity(ctx context.Context, name string, channel string) (string, error) {
	return NewSlackActivityContext().postMessageActivityImpl(ctx, name, channel)
}

func (c *SlackActivityContext) postMessageActivityImpl(ctx context.Context, name string, channel string) (string, error) {
	instrumentation.Log(fmt.Sprintf("SlackMessageActivity name:%s", name))

	stackTrace := "Traceback (most recent call last):\n  File \"tb.py\", line 15, in <module>\n    a()\n  File \"tb.py\", line 3, in a\n    j = b(i)\n  File \"tb.py\", line 9, in b\n    c()\n  File \"tb.py\", line 13, in c\n    error()\nNameError: name 'error' is not defined\n"
	attachment := slack.Attachment{
		Pretext: "Does this look like an error?",
		Text:    stackTrace,
	}
	channelID, timestamp, err := c.client.PostMessage(channel, slack.MsgOptionText(name, false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		instrumentation.Log(fmt.Sprintf("Error:%s", err.Error()))
		return "", err
	}

	instrumentation.Log(fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))
	return "", nil
}
