package gateway

import (
	"context"
)

type Identifier struct {
	IdentifyCommand
}

func DefaultIdentifier(token string) Identifier {
	return NewIdentifier(DefaultIdentifyCommand(token))
}

func NewIdentifier(data IdentifyCommand) Identifier {
	return Identifier{
		IdentifyCommand: data,
	}
}

func (id *Identifier) QueryGateway(ctx context.Context) (gatewayURL string, err error) {
	return URL(ctx, id.Token)
}

func (id *Identifier) Wait(ctx context.Context) error {
	return nil
}

func DefaultIdentifyCommand(token string) IdentifyCommand {
	return IdentifyCommand{
		Token: token,
	}
}

func (i *IdentifyCommand) AddIntents(intents Intents) {
	i.Intents |= uint(intents)
}
