package notify

import (
	"context"
	"dga-cp-lab-kafka/app/entities"
)

func (useCase *useCase) ProducerNewMessage(input *entities.Message) {
	ctx := context.Background()
	go func(ctx context.Context) {
		_ = useCase.NotifyRepo.ProducerNewMessage(input)
	}(ctx)
}
