package notify

import (
	"dga-cp-lab-kafka/app/entities"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type repo struct {
	Producer *kafka.Producer
}

type Repo interface {
	ProducerNewMessage(input *entities.Message) error
}

func InitRepo(producer *kafka.Producer) Repo {
	return &repo{
		Producer: producer,
	}
}
