package bot

import (
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

type Context struct {
	*Bot
	Event ws.Event

	stop bool
}

func (c *Context) Next() {
	c.stop = false
}
