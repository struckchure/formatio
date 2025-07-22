package lib

import (
	"context"
	"log"

	"github.com/ably/ably-go/ably"
)

type AblyInterface struct {
	client *ably.Realtime
}

type PublishAblyArgs struct {
	Event   string
	Content interface{}
}

type SubscribeAblyArgs struct {
	Event    string
	Callback func(content interface{}) error
}

func (a *AblyInterface) Publish(args PublishAblyArgs) error {
	// TODO: publish only if there's someone listening
	channel := a.client.Channels.Get(args.Event)

	err := channel.PublishAsync("", args.Content, func(err error) {})
	if err != nil {
		log.Println(err)

		return err
	}

	return nil
}

func (a *AblyInterface) Subscribe(args SubscribeAblyArgs) error {
	channel := a.client.Channels.Get(args.Event)
	_, err := channel.Subscribe(
		context.Background(),
		"",
		func(msg *ably.Message) {
			args.Callback(msg.Data)
		},
	)
	if err != nil {
		log.Println(err)

		return err
	}

	return nil
}

func NewAbly(client *ably.Realtime) AblyInterface {
	return AblyInterface{client: client}
}

func NewAblyConnection(env Env) (*ably.Realtime, error) {
	// Connect to Ably with your API key
	client, err := ably.NewRealtime(ably.WithKey(env.ABLY_API_KEY), ably.WithAutoConnect(false))
	if err != nil {
		log.Println("ably: ", err)
	}
	client.Connection.OnAll(func(change ably.ConnectionStateChange) {
		log.Printf("ably: event=%s state=%s reason=%s", change.Event, change.Current, change.Reason)
	})
	client.Connect()

	return client, err
}
