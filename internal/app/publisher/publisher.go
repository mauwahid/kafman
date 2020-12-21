package publisher

import (
	"github.com/json-iterator/go"
	"github.com/mauwahid/kafman/internal/domain/publisher"
	"github.com/mauwahid/kafman/internal/infra/config"
	"github.com/mauwahid/kafman/internal/infra/errs"
	"github.com/mauwahid/kafman/internal/infra/queue/producer"
	"github.com/mauwahid/kafman/internal/presenter/dto"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Publisher struct {
	pub   publisher.Publisher
	queue publisher.Queue
}

func New() Publisher {

	q, err := producer.New(config.Get().GetString("producer.brokers"))

	if err != nil {
		panic(err)
	}

	return Publisher{pub: publisher.NewPublisher(), queue: q}
}

func (p Publisher) Publish(pubReq dto.PubRequest) (pubRes dto.PubResponse, err error) {

	var data []byte
	if data, err = json.Marshal(pubReq.Message); err != nil {
		pubRes = errs.Error(data, err)
		return
	}

	var key string
	if key, err = p.pub.Publish(pubReq.Topic, string(data), p.queue); err != nil {
		pubRes = errs.Error(data, err)
		return
	}

	return errs.Success(key, string(data)), nil
}
