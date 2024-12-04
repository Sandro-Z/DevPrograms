package auditlog

import (
	"git.ana/xjtuana/qqguildgo/bot"
)

var plugin = NewPlugin("auditlog")

func init() { bot.Add(plugin.Name, plugin.Handler) }

type PluginAuditLog struct {
	Name string
}

func NewPlugin(name string) *PluginAuditLog {
	return &PluginAuditLog{
		Name: name,
	}
}

func Middleware(c *bot.Context) {
	c.Next()
}

func (p *PluginAuditLog) Handler(c *bot.Context) {
	// c.SendText("pong")
}
