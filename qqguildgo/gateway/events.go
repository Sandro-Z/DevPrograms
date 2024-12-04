package gateway

import (
	"regexp"
	"strings"

	"git.ana/xjtuana/qqguildgo/qqguild"
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

func init() {
	OpUnmarshalers.Add(
		func() ws.Event { return new(HeartbeatAckEvent) },
		func() ws.Event { return new(MessageCreateEvent) },
		func() ws.Event { return new(AtMessageCreateEvent) },
		func() ws.Event { return new(DirectMessageCreateEvent) },
		func() ws.Event { return new(ReadyEvent) },
		func() ws.Event { return new(ResumedEvent) },
		func() ws.Event { return new(HeartbeatCommand) },
		func() ws.Event { return new(HelloEvent) },
		func() ws.Event { return new(ResumeCommand) },
	)
}

type Event ws.Event

const (
	dispatchOp        ws.OpCode = 0  // recv
	heartbeatOp       ws.OpCode = 1  // send/recv
	identifyOp        ws.OpCode = 2  // send
	resumeOp          ws.OpCode = 6  // send
	reconnectOp       ws.OpCode = 7  // recv
	invalidSessionOp  ws.OpCode = 9  // recv
	helloOp           ws.OpCode = 10 // recv
	heartbeatAckOp    ws.OpCode = 11 // recv/reply
	httpCallbackAckOp ws.OpCode = 12 // reply
)

var OpUnmarshalers = ws.NewOpUnmarshalers()

type ReadyEvent struct {
	Version   int          `json:"version"`
	SessionID string       `json:"session_id"`
	User      qqguild.User `json:"user"`
}

func (*ReadyEvent) Op() ws.OpCode           { return dispatchOp }
func (*ReadyEvent) EventType() ws.EventType { return "READY" }

type ResumedEvent struct{}

func (*ResumedEvent) Op() ws.OpCode           { return dispatchOp }
func (*ResumedEvent) EventType() ws.EventType { return "RESUMED" }

type MessageCreateEvent struct {
	qqguild.Message
	Member       *qqguild.Member `json:"member,omitempty"`
	cmd, content string
}

func (*MessageCreateEvent) Op() ws.OpCode           { return dispatchOp }
func (*MessageCreateEvent) EventType() ws.EventType { return "MESSAGE_CREATE" }

func (e *MessageCreateEvent) EqualFolds(cmds ...string) bool {
	if len(cmds) == 0 {
		return false
	}
	_ = e.GetCommand()
	for _, v := range cmds {
		if strings.EqualFold(e.cmd, v) {
			return true
		}
	}
	return false
}

func (e *MessageCreateEvent) GetCommand() string {
	if e.cmd == "" {
		content := regexp.MustCompile(`<@!\d+>`).ReplaceAllString(e.Content, "")
		e.cmd, e.content, _ = strings.Cut(strings.Trim(content, " \u00A0/"), " ")
	}
	return e.cmd
}

func (e *MessageCreateEvent) GetContent() string {
	return e.content
}

type AtMessageCreateEvent MessageCreateEvent

func (*AtMessageCreateEvent) Op() ws.OpCode           { return dispatchOp }
func (*AtMessageCreateEvent) EventType() ws.EventType { return "AT_MESSAGE_CREATE" }

type DirectMessageCreateEvent MessageCreateEvent

func (*DirectMessageCreateEvent) Op() ws.OpCode           { return dispatchOp }
func (*DirectMessageCreateEvent) EventType() ws.EventType { return "DIRECT_MESSAGE_CREATE" }

type HeartbeatCommand int

func (*HeartbeatCommand) Op() ws.OpCode           { return heartbeatOp }
func (*HeartbeatCommand) EventType() ws.EventType { return "" }

type HeartbeatAckEvent struct{}

func (*HeartbeatAckEvent) Op() ws.OpCode           { return heartbeatAckOp }
func (*HeartbeatAckEvent) EventType() ws.EventType { return "" }

type IdentifyCommand struct {
	Token   string `json:"token"`
	Intents uint   `json:"intents"`
}

func (*IdentifyCommand) Op() ws.OpCode           { return identifyOp }
func (*IdentifyCommand) EventType() ws.EventType { return "" }

type ReconnectEvent struct{}

func (*ReconnectEvent) Op() ws.OpCode           { return reconnectOp }
func (*ReconnectEvent) EventType() ws.EventType { return "" }

type ResumeCommand struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"seq"`
}

func (*ResumeCommand) Op() ws.OpCode           { return resumeOp }
func (*ResumeCommand) EventType() ws.EventType { return "" }

type HelloEvent struct {
	HeartbeatInterval qqguild.Milliseconds `json:"heartbeat_interval"`
}

func (*HelloEvent) Op() ws.OpCode           { return helloOp }
func (*HelloEvent) EventType() ws.EventType { return "" }
