package entities

import (
	"time"
)

type Message struct {
	Topic     string
	Key       string
	Value     string
	Timestamp time.Time
}
