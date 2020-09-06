package consumers

import (
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rjar2020/post-delivery/env"
)

//StartConsumer initializes and runs a kafka consumer
func StartConsumer(topic string, groupID string, wg *sync.WaitGroup) error {
	consumer, err := createConsumer(groupID)
	if err != nil {
		log.Printf("Error initializaing consumer for topic %s - Group Id: %s", topic, groupID)
		return err
	}
	log.Printf("Consuming from for topic %s - Group Id: %s", topic, groupID)
	subscribeAndRunConsumer(topic, groupID, consumer)
	consumer.Close()
	defer wg.Done()
	return err
}

func createConsumer(groupID string) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv(env.KafkaBrokers),
		"group.id":          groupID,
		"auto.offset.reset": os.Getenv(env.KafkaConsumerOffsetReset)})
}

func subscribeAndRunConsumer(topic string, groupID string, consumer *kafka.Consumer) error {
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err == nil {
		for true {
			ev := consumer.Poll(2000)
			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("Message on %s: %s",
					e.TopicPartition, string(e.Value))
			case kafka.PartitionEOF:
				log.Printf("Reached %v", e)
			case kafka.Error:
				log.Panic("Error when reading from kafka", e)
			default:
				log.Printf("Topic: %v - GroupId: %#v. No message during last poll %v\n", topic, groupID, e)
			}
		}
	} else {
		log.Printf("Error subscribing to topics: %v", topic)
	}
	return err
}
