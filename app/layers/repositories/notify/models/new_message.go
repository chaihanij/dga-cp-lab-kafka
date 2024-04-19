package models

import (
	"dga-cp-lab-kafka/app/constants"
	"dga-cp-lab-kafka/app/entities"
	"github.com/google/uuid"
	"time"
)

func NewMessage(input *entities.Message) *Message {
	return &Message{
		Topic:     constants.TopicNameDgaCpoLabTeamTopic,
		Key:       uuid.New().String(),
		Value:     input.Value,
		Timestamp: time.Now(),
	}
}
