package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"git.ana/xjtuana/qqguildgo/api"
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

const CodeInvalidSequence = 4007

type State struct {
	Identifier Identifier
	SessionID  string
	Sequence   int64
}

type Gateway struct {
	gateway *ws.Gateway
	state   State

	beatMutex  sync.Mutex
	sentBeat   time.Time
	echoBeat   time.Time
	retryTimer time.Timer
}

func URL(ctx context.Context, token string) (string, error) {
	return api.NewClient(token).GatewayURL()
}

func NewWithIntents(ctx context.Context, token string, intents ...Intents) (*Gateway, error) {
	var allIntents Intents
	for _, intent := range intents {
		allIntents |= intent
	}
	g, err := New(ctx, token)
	if err != nil {
		return nil, err
	}
	g.AddIntents(allIntents)
	return g, nil
}

func New(ctx context.Context, token string) (*Gateway, error) {
	return NewWithIdentifier(ctx, DefaultIdentifier(token))
}

func NewWithIdentifier(ctx context.Context, id Identifier) (*Gateway, error) {
	gatewayURL, err := id.QueryGateway(ctx)
	if err != nil {
		return nil, err
	}
	gateway := NewCustomWithIdentifier(gatewayURL, id, nil)
	return gateway, nil
}

func NewCustomWithIdentifier(gatewayURL string, id Identifier, opts *ws.GatewayOpts) *Gateway {
	return NewFromState(gatewayURL, State{Identifier: id}, opts)
}

var DefaultGatewayOpts = ws.GatewayOpts{
	ReconnectDelay: func(try int) time.Duration {
		// minimum 4 seconds
		return time.Duration(4+(2*try)) * time.Second
	},
	// FatalCloseCodes contains the default gateway close codes that will cause
	// the gateway to exit. In other words, it's a list of unrecoverable close
	// codes.
	FatalCloseCodes: []int{
		4004, // authentication failed
		4010, // invalid shard sent
		4011, // sharding required
		4012, // invalid API version
		4013, // invalid intents
		4014, // disallowed intents
	},
	DialTimeout:           0,
	ReconnectAttempt:      0,
	AlwaysCloseGracefully: true,
}

func NewFromState(gatewayURL string, state State, opts *ws.GatewayOpts) *Gateway {
	if opts == nil {
		opts = &DefaultGatewayOpts
	}
	gw := ws.NewGateway(ws.NewWebsocket(ws.NewCodec(OpUnmarshalers), gatewayURL), opts)
	return &Gateway{
		gateway: gw,
		state:   state,
	}
}

func (g *Gateway) AddIntents(i Intents) {
	g.gateway.AssertIsNotRunning()
	g.state.Identifier.AddIntents(i)
}

func (g *Gateway) LastError() error {
	return g.gateway.LastError()
}

func (g *Gateway) Connect(ctx context.Context) <-chan ws.Op {
	return g.gateway.Connect(ctx, &gatewayImpl{Gateway: g})
}

type gatewayImpl struct {
	*Gateway
	lastSentBeat time.Time
}

func (g *gatewayImpl) invalidate() {
	g.state.SessionID = ""
	g.state.Sequence = 0
}

func (g *gatewayImpl) sendIdentify(ctx context.Context) error {
	if err := g.state.Identifier.Wait(ctx); err != nil {
		return fmt.Errorf("can't wait for identify(): %w", err)
	}
	return g.gateway.Send(ctx, &g.state.Identifier.IdentifyCommand)
}

func (g *gatewayImpl) sendResume(ctx context.Context) error {
	return g.gateway.Send(ctx, &ResumeCommand{
		Token:     g.state.Identifier.Token,
		SessionID: g.state.SessionID,
		Sequence:  g.state.Sequence,
	})
}

func (g *gatewayImpl) OnOp(ctx context.Context, op ws.Op) bool {
	p, _ := json.Marshal(op)
	log.Println("gateway op:", string(p))

	if op.Code == dispatchOp {
		g.state.Sequence = op.Sequence
	}

	switch data := op.Data.(type) {
	case *ws.CloseEvent:
		if data.Code == CodeInvalidSequence {
			// Invalid sequence.
			g.invalidate()
		}

		g.gateway.QueueReconnect()

	case *HelloEvent:
		g.gateway.ResetHeartbeat(data.HeartbeatInterval.Duration())

		// Send Discord either the Identify packet (if it's a fresh
		// connection), or a Resume packet (if it's a dead connection).
		if g.state.SessionID == "" || g.state.Sequence == 0 {
			// SessionID is empty, so this is a completely new session.
			if err := g.sendIdentify(ctx); err != nil {
				g.gateway.SendErrorWrap(err, "failed to send identify")
				g.gateway.QueueReconnect()
			}
		} else {
			if err := g.sendResume(ctx); err != nil {
				g.gateway.SendErrorWrap(err, "failed to send resume")
				g.gateway.QueueReconnect()
			}
		}

	case *HeartbeatCommand:
		g.SendHeartbeat(ctx)

	case *HeartbeatAckEvent:
		now := time.Now()

		g.beatMutex.Lock()
		g.sentBeat = g.lastSentBeat
		g.echoBeat = now
		g.beatMutex.Unlock()

	case *ReconnectEvent:
		g.gateway.QueueReconnect()

	case *ReadyEvent:
		g.state.SessionID = data.SessionID
	}

	return true
}

// SendHeartbeat sends a heartbeat with the gateway's current sequence.
func (g *gatewayImpl) SendHeartbeat(ctx context.Context) {
	g.lastSentBeat = time.Now()
	sequence := HeartbeatCommand(g.state.Sequence)
	if err := g.gateway.Send(ctx, &sequence); err != nil {
		g.gateway.SendErrorWrap(err, "heartbeat error")
		g.gateway.QueueReconnect()
	}
}

// Close closes the state.
func (g *gatewayImpl) Close() error {
	g.retryTimer.Stop()
	g.invalidate()
	return nil
}
