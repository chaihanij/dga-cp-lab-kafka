package models

import (
	"dga-cp-lab-kafka/app/entities"
	"github.com/jinzhu/copier"
	"time"
)

type Message struct {
	Topic     string    `json:"topic"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func (value *Message) ToEntity() (*entities.Message, error) {
	var result entities.Message
	err := copier.Copy(&result, value)
	return &result, err
}
