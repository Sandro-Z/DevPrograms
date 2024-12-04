package bot

import (
	"context"
	"fmt"
	"sync"

	"git.ana/xjtuana/qqguildgo/api"
	"git.ana/xjtuana/qqguildgo/gateway"
	"github.com/diamondburned/arikawa/v3/utils/handler"
	"github.com/diamondburned/arikawa/v3/utils/ws"
	"github.com/diamondburned/arikawa/v3/utils/ws/ophandler"
)

var (
	this = New("")

	Add = this.Add
	Use = this.Use
	Run = this.Run
)

func NewClient(token string) {
	this.Client = api.NewClient(token)
	this.state.id = gateway.DefaultIdentifier(token)
	this.state.id.AddIntents(gateway.IntentsAll)
}

func SetClient(client *api.Client) {
	this.Client = client
}

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Bot struct {
	*api.Client
	*handler.Handler

	middleware HandlersChain
	handlers   map[string]HandlersChain

	state *state
}

type state struct {
	sync.Mutex
	id      gateway.Identifier
	gateway *gateway.Gateway

	ctx    context.Context
	cancel context.CancelFunc
	doneCh <-chan struct{}
}

func New(token string) *Bot {
	b := &Bot{
		Client:     api.NewClient(token),
		Handler:    handler.New(),
		middleware: HandlersChain{},
		handlers:   map[string]HandlersChain{},
		state: &state{
			id: gateway.DefaultIdentifier(token),
		},
	}
	b.AddSyncHandler(b.handleAtMessageCreateEvent)
	b.AddSyncHandler(b.handleDirectMessageCreateEvent)
	b.AddSyncHandler(b.handleMessageCreateEvent)
	return b
}

func NewWithIntents(token string, intents ...gateway.Intents) *Bot {
	b := New(token)
	var allIntent gateway.Intents
	for _, intent := range intents {
		allIntent |= intent
	}
	b.state.id.AddIntents(allIntent)
	return b
}

func (b *Bot) NewContext(e ws.Event) *Context {
	return &Context{
		Bot:   b,
		Event: e,
	}
}

func (b *Bot) Add(name string, handlers ...HandlerFunc) *Bot {
	if _, ok := b.handlers[name]; ok {
		b.handlers[name] = append(b.handlers[name], handlers...)
	} else {
		b.handlers[name] = handlers
	}
	return b
}

func (b *Bot) Use(middleware ...HandlerFunc) *Bot {
	b.middleware = append(b.middleware, middleware...)
	return b
}

func (b *Bot) Run(ctx context.Context) error {
	evCh := make(chan any)

	b.state.Lock()
	defer b.state.Unlock()

	if b.state.cancel != nil {
		if err := b.close(ctx); err != nil {
			return err
		}
	}

	if b.state.gateway == nil {
		g, err := gateway.NewWithIdentifier(ctx, b.state.id)
		if err != nil {
			return err
		}
		b.state.gateway = g
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.state.ctx = ctx
	b.state.cancel = cancel

	// TODO: change this to AddSyncHandler.
	rm := b.AddHandler(evCh)
	defer rm()

	opCh := b.state.gateway.Connect(ctx)
	b.state.doneCh = ophandler.Loop(opCh, b.Handler)

	for {
		select {
		case <-ctx.Done():
			b.close(ctx)
			return ctx.Err()

		case <-b.state.doneCh:
			// Event loop died.
			return b.state.gateway.LastError()

		case ev := <-evCh:
			switch ev.(type) {
			case *gateway.ReadyEvent, *gateway.ResumedEvent:
				return nil
			}
		}
	}
}

func (b *Bot) Close() error {
	b.state.Lock()
	defer b.state.Unlock()

	return b.close(context.Background())
}

func (b *Bot) close(ctx context.Context) error {
	if b.state.cancel == nil {
		return fmt.Errorf("Session is already closed")
	}

	b.state.cancel()
	b.state.cancel = nil
	b.state.ctx = nil

	// Wait until we've successfully disconnected.
	select {
	case <-ctx.Done():
		return fmt.Errorf("cannot wait for gateway exit %s", ctx.Err())
	case <-b.state.doneCh:
		// ok
	}

	b.state.doneCh = nil

	return b.state.gateway.LastError()
}
