package api

import (
	"git.ana/xjtuana/qqguildgo/qqguild"
)

type ChannelThreads struct {
	Threads  []qqguild.Thread `json:"threads"`
	Finished int64            `json:"is_finish"`
}

func (c *Client) Threads(channelID qqguild.ChannelID) ([]qqguild.Thread, error) {
	var cts ChannelThreads
	return cts.Threads, c.GetJSON(EndpointChannels+channelID.String()+"/threads", &cts)
}

func (c *Client) Posts(channelID qqguild.ChannelID, threadID string) ([]qqguild.Post, error) {
	var ps []qqguild.Post
	return ps, c.GetJSON(EndpointChannels+channelID.String()+"/threads/"+threadID, &ps)
}
