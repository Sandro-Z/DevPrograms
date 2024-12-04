package bot

import (
	"git.ana/xjtuana/qqguildgo/gateway"
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

func (b *Bot) handleMessageCreateEvent(e *gateway.MessageCreateEvent) {
	b.handle(e)
}

func (b *Bot) handleAtMessageCreateEvent(e *gateway.AtMessageCreateEvent) {
	b.handle(e)
}

func (b *Bot) handleDirectMessageCreateEvent(e *gateway.DirectMessageCreateEvent) {
	b.handle(e)
}

func (b *Bot) handle(e ws.Event) {
	c := b.NewContext(e)
	for _, f := range b.middleware {
		c.stop = true
		if f(c); c.stop {
			return
		}
	}
	for _, fs := range b.handlers {
		for _, f := range fs {
			go f(c)
		}
	}
}
