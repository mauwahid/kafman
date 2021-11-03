package process

import "github.com/mauwahid/kafman/internal/app/subscriber"

type Subscriber struct{}

func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

func (s *Subscriber) Run() {
	app := subscriber.New()
	app.Subscribe()
}
