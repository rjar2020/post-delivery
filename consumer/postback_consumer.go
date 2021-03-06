package consumer

import (
	"log"
	"os"
	"sync"

	"github.com/rjar2020/post-delivery/producer"

	"github.com/rjar2020/post-delivery/model"
	"github.com/rjar2020/post-delivery/service"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rjar2020/post-delivery/env"
)

//I know this is not needed, I just wanted to try iota
//Besides, this is not prod code, is just a test and this is a Golang feature to try and showcase.
const (
	//OneSecond for kafka consumer polling
	OneSecond = (iota + 1) * 1000
	//TwoSeconds for kafka consumer polling
	TwoSeconds = (iota + 1) * 1000
)

//StartPostBackConsumer initializes and runs a kafka consumer
func StartPostBackConsumer(topic string, groupID string, wg *sync.WaitGroup) error {
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
	if err != nil {
		log.Printf("Error subscribing to topics: %v", topic)
		return err
	}

	for true {
		ev := consumer.Poll(TwoSeconds)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("Message on %s: %s", e.TopicPartition, string(e.Value))
			processPostBack(e)
		case kafka.PartitionEOF:
			log.Printf("Reached %v", e)
		case kafka.Error:
			log.Printf("Error when reading from kafka %v", e)
		default:
			log.Printf("Topic: %v - GroupId: %#v. No message during last poll %v\n", topic, groupID, e)
		}
	}

	return err
}

func processPostBack(message *kafka.Message) {
	postBack, err := model.FromJSONtoPostback(message.Value)
	if err != nil {
		log.Printf("Error decoding kafka message: %s", err)
		return
	}

	log.Printf("Decoded message on %s: %#v", message.TopicPartition, postBack)
	url := service.ToURL(postBack)
	_, err = service.DeliverPostback(postBack.Endpoint.Method, url)
	if err != nil {
		log.Printf("Error when processing postback. Sending it to dead letter topic: %v", err)
		producer.Produce(string(message.Value), os.Getenv(env.KafkaDeadPostBackTopic))
	}
}
