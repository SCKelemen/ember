package actor

import (
	"fmt"

	"github.com/sckelemen/ember/src/ember/event"
)

type Actor interface {
	Listen(address string) error
	Dispatch(event.Event) error
	Handle(event.Event) error
}

type SnowflakeActor struct {
	pid      int64 // snowflake
	sid      int64 // sub id
	events   chan event.Event
	sequence int64
}

func (actor *SnowflakeActor) Listen(address string) error {
	actor.events = make(chan event.Event, 1024)

	for event := range actor.events {
		if err := actor.Handle(event); err != nil {
			return err
		}

	}

	return nil
}

func (actor *SnowflakeActor) Handle(event event.Event) error {
	switch event.Type {
	case event.Hello:
		return actor.onHelloEvent(event)
	default:
		return fmt.Errorf("no handler identified for event %s", event.Type)
	}
}

func (actor *SnowflakeActor) Dispatch(event event.Event) error {
	actor.events <- event
	return nil
}

// it would be really great if go supported yielding enumberables...

type SFActor struct {
}
