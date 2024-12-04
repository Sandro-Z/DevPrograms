package bot

import (
	"git.ana/xjtuana/qqguildgo/qqguild"
)

func (b *Bot) SendMessage(channelID qqguild.ChannelID, messageID qqguild.MessageID, content string, embeds ...qqguild.Embed) error {
	msg, err := b.Client.SendMessage(channelID, messageID, content, embeds...)
	if err != nil {
		return err
	}
	_ = msg
	return nil
}
