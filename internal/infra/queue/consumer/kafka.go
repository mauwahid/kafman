package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"strings"
)

type Consumer struct {
	subscriber *kafka.Subscriber
	group      string
	brokers    string
	consume    Consume
}

type Consume func(topic, key string, message []byte)

func New(group, brokers string, consume Consume) Consumer {
	return Consumer{
		group:   group,
		brokers: brokers,
		consume: consume,
	}
}

func (c *Consumer) Subscribe(topics []string) error {

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	kBrokers := strings.Split(c.brokers, ",")
	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               kBrokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         c.group,
		},
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		return err
	}

	c.subscriber = subscriber

	for i := range topics {
		messages, err := subscriber.Subscribe(context.Background(), topics[i])
		if err != nil {
			return err
		}

		go c.process(topics[i], messages)
	}

	return nil
}

func (c Consumer) process(topic string, messages <-chan *message.Message) {
	for msg := range messages {
		c.consume(topic, msg.UUID, msg.Payload)
		msg.Ack()
	}
}
