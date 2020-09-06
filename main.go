package main

import (
	"log"
	"sync"

	"github.com/rjar2020/post-delivery/env"

	"github.com/rjar2020/post-delivery/consumers"
	"github.com/rjar2020/post-delivery/producers"
)

func main() {
	log.Printf("####### Strating post delivery application #######")
	env.LoadEnv()
	var wg sync.WaitGroup
	wg.Add(1)
	topic := "test-topic"
	producers.Produce("Hello world before consuming", topic)
	go consumers.StartConsumer(topic, "postback-processor", &wg)
	producers.Produce("Hello world after starting consumer", topic)
	wg.Wait()
}
