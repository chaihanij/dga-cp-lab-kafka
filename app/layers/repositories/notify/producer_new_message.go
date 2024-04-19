package notify

import (
	"dga-cp-lab-kafka/app/entities"
	"dga-cp-lab-kafka/app/layers/repositories/notify/models"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

func (r *repo) ProducerNewMessage(input *entities.Message) error {
	payload := models.NewMessage(input)
	value, _ := json.Marshal(payload)
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &payload.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(payload.Key),
		Value:          value,
	}
	log.Printf("Producing message: %v", message)
	err := r.Producer.Produce(message, nil)
	if err != nil {

		log.Printf("Failed to produce message: %v", err)
		return err
	}
	return nil
}
