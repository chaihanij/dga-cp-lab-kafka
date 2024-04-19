package notify

import (
	"dga-cp-lab-kafka/app/entities"
	notifyRepo "dga-cp-lab-kafka/app/layers/repositories/notify"
)

type useCase struct {
	NotifyRepo notifyRepo.Repo
}

type UseCase interface {
	ProducerNewMessage(input *entities.Message)
}

func InitUseCase(notifyRepo notifyRepo.Repo) UseCase {
	return &useCase{
		NotifyRepo: notifyRepo,
	}
}
