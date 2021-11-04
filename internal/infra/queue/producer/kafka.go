package producer

import (
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Producer struct {
	publisher *kafka.Publisher
}

func New(brokers string) (*Producer, error) {

	kBrokers := strings.Split(brokers, ",")
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   kBrokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(true, true),
	)

	if err != nil {
		return &Producer{}, err
	}

	producer := Producer{
		publisher: publisher,
	}

	return &producer, nil
}

func (p *Producer) Send(topic string, data []byte) (string, error) {
	uuid := watermill.NewUUID()
	msg := message.NewMessage(watermill.NewUUID(), data)
	if err := p.publisher.Publish(topic, msg); err != nil {
		return uuid, err
	}
	return uuid, nil
}
