package producer

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rjar2020/post-delivery/env"
)

//Produce a message to kafka
func Produce(message string, topic string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv(env.KafkaBrokers),
		"client.id":         "postback-producer",
		"acks":              os.Getenv(env.KafkaProducersAck)})
	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
	} else {
		err = produceToTopic(p, message, topic)
	}
	return err
}

func produceToTopic(producer *kafka.Producer, message string, topic string) error {
	deliveryChan := make(chan kafka.Event, 10000)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message)},
		deliveryChan,
	)
	if err != nil {
		log.Printf("Delivery failed: %v\n", err)
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(deliveryChan)
	return m.TopicPartition.Error
}
