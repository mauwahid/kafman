package publisher

type Publisher struct{}

func NewPublisher() Publisher {
	return Publisher{}
}

func (p *Publisher) Publish(topic, message string, q Queue) (string, error) {
	return q.Send(topic, message)
}
