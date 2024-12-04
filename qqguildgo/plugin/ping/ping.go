package ping

import (
	"log"

	"git.ana/xjtuana/qqguildgo/bot"
	"git.ana/xjtuana/qqguildgo/gateway"
)

var plugin = NewPlugin("ping")

func init() { bot.Add(plugin.Name, plugin.Handler) }

type PluginPing struct {
	Name string
}

func NewPlugin(name string) *PluginPing {
	return &PluginPing{
		Name: name,
	}
}

func (p *PluginPing) Handler(c *bot.Context) {
	var e gateway.MessageCreateEvent
	switch v := c.Event.(type) {
	default:
		return
	case *gateway.AtMessageCreateEvent:
		e = gateway.MessageCreateEvent(*v)
	case *gateway.DirectMessageCreateEvent:
		e = gateway.MessageCreateEvent(*v)
	}
	if !e.EqualFolds("PING", "PINGPONG", "PING-PONG") {
		return
	}
	p.handleEvent(c, &e)
}

func (p *PluginPing) handleEvent(c *bot.Context, e *gateway.MessageCreateEvent) {
	content := "pong!"
	if !e.DirectMessage {
		content = e.Author.Mention() + " " + content
	}
	if len(e.GetContent()) > 0 {
		content += " " + e.GetContent()
	}
	if err := c.SendMessage(e.ChannelID, e.ID, content); err != nil {
		log.Println(err)
	}
}
