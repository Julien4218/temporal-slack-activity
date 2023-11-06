package models

type SlackActivityData struct {
	ChannelId            string
	FirstResponseWarning string
	Attachment           MessageAttachment
}
