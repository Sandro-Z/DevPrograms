package api

import (
	"git.ana/xjtuana/qqguildgo/qqguild"
)

type SendMessageData struct {
	Content string          `json:"content,omitempty"`
	Embeds  []qqguild.Embed `json:"embeds,omitempty"`
	Photo   qqguild.URL     `json:"image,omitempty"`

	MessageID qqguild.MessageID         `json:"msg_id,omitempty"`
	Reference *qqguild.MessageReference `json:"message_reference,omitempty"`
}

func (c *Client) SendMessage(channelID qqguild.ChannelID, messageID qqguild.MessageID, content string, embeds ...qqguild.Embed) (*qqguild.Message, error) {
	var msg *qqguild.Message
	return msg, c.PostJSON(EndpointChannels+channelID.String()+"/messages", SendMessageData{
		MessageID: messageID,
		Content:   content,
		Embeds:    embeds,
	}, &msg)
}

func (c *Client) DeleteMessage(channelID qqguild.ChannelID, messageID qqguild.MessageID, hide ...bool) error {
	params := "hidetip=false"
	if len(hide) > 0 && hide[0] {
		params = "hidetip=true"
	}
	return c.DeleteJSON(EndpointChannels+channelID.String()+"/messages/"+messageID+"?"+params, nil)
}
