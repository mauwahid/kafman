package subscriber

import (
	"github.com/google/martian/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/mauwahid/kafman/internal/domain/subscriber"
	"github.com/mauwahid/kafman/internal/infra/config"
	"github.com/mauwahid/kafman/internal/infra/queue/consumer"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Subscriber struct {
	sub   subscriber.Subscriber
	queue subscriber.Queue
}

func New() Subscriber {

	group := config.Get().GetString("consumer.group")
	brokers := config.Get().GetString("consumer.brokers")
	topics := config.Get().GetString("consumer.topics")

	sub := subscriber.NewSubsriber(strings.Split(topics, ","))
	q := consumer.New(group, brokers, sub.ConsumeMessage)
	return Subscriber{sub: sub, queue: &q}
}

func (s Subscriber) Subscribe() {
	if err := s.sub.Subscribe(s.queue); err != nil {
		log.Errorf("err %s", err.Error())
	}
}
