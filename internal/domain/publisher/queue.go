package publisher

type Queue interface {
	Send(topic string, message []byte) (string, error)
}
