package antispam

import (
	"log"

	"git.ana/xjtuana/qqguildgo/bot"
	"git.ana/xjtuana/qqguildgo/gateway"
)

var plugin = NewPlugin("antispam")

func init() { bot.Add(plugin.Name, plugin.Handler) }

type PluginAntiSpam struct {
	Name string
}

func NewPlugin(name string) *PluginAntiSpam {
	return &PluginAntiSpam{
		Name: name,
	}
}

func Middleware(c *bot.Context) {
	e, ok := c.Event.(*gateway.MessageCreateEvent)
	if !ok {
		return
	}
	if err := checkSpam(c, e); err != nil {
		log.Println(err)
		return
	}
	c.Next()
}

func checkSpam(c *bot.Context, e *gateway.MessageCreateEvent) error {
	return nil
}

func (p *PluginAntiSpam) Handler(c *bot.Context) {
	// c.SendText("pong")
}
