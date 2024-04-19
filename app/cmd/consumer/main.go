package main

import (
	"dga-cp-lab-kafka/app/constants"
	"dga-cp-lab-kafka/app/database"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

func main() {

	// creates a new consumer and subscribes to your topic
	consumer := database.KafkaInitConsumer()
	err := consumer.SubscribeTopics([]string{constants.TopicNameDgaCpoLabTeamTopic}, nil)
	if err != nil {
		log.Panicf("Failed to subscribe to topic: %v\n", err)
	}

	run := true
	for run {
		// consumes messages from the subscribed topic and prints them to the console
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			log.Printf("Consumed event from topic %s: key = %-10s value = %s\n", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		case kafka.Error:
			log.Printf("Error: %v\n", ev)
			run = false
		}
	}

	// closes the consumer connection
	err = consumer.Close()
	if err != nil {
		return
	}
}
