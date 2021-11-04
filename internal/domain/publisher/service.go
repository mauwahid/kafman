package publisher

type Publisher struct{}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(topic string, message []byte, q Queue) (string, error) {
	return q.Send(topic, message)
}
