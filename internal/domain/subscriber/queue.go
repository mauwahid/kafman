package subscriber

type Queue interface {
	Subscribe(topics []string) error
}
