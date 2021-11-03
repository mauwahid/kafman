package subscriber

import (
	"fmt"
)

type Subscriber struct {
	topics []string
}

func NewSubsriber(topics []string) *Subscriber {
	return &Subscriber{topics: topics}
}

func (s *Subscriber) Subscribe(q Queue) error {
	return q.Subscribe(s.topics)
}

func (s *Subscriber) ConsumeMessage(topic, key string, message []byte) {
	fmt.Printf("message consumed with key : %s, topic : %s, message : %s \n", key, topic, string(message))
}
