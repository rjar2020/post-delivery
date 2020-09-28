package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rjar2020/post-delivery/consumer"
	"github.com/rjar2020/post-delivery/controller"
	"github.com/rjar2020/post-delivery/env"
)

func main() {
	log.Printf("####### Strating post delivery application #######")
	env.LoadEnv()
	controller.RegisterControllers()
	var wg sync.WaitGroup
	wg.Add(8)
	httpPort := ":" + os.Getenv(env.APIPort)
	topic := os.Getenv(env.KafkaPostBackTopic)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	http.ListenAndServe(httpPort, nil)
	wg.Wait()
}
