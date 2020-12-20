package publisher

type Queue interface {
	Send(topic, message string) (string, error)
}
