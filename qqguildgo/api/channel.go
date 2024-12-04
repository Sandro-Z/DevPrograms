package api

import (
	"git.ana/xjtuana/qqguildgo/qqguild"
)

var EndpointChannels = Endpoint + "channels/"

func (c *Client) Channels(guildID qqguild.GuildID) ([]qqguild.Channel, error) {
	var chs []qqguild.Channel
	return chs, c.GetJSON(EndpointGuilds+guildID.String()+"/channels", &chs)
}
