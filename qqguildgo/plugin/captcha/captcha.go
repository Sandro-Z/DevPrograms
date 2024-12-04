package captcha

import (
	"git.ana/xjtuana/qqguildgo/bot"
)

var plugin = NewPlugin("captcha")

func init() { bot.Add(plugin.Name, plugin.Handler) }

type PluginCaptcha struct {
	Name string
}

func NewPlugin(name string) *PluginCaptcha {
	return &PluginCaptcha{
		Name: name,
	}
}

func Middleware(c *bot.Context) {
	c.Next()
}

func (p *PluginCaptcha) Handler(c *bot.Context) {
}
