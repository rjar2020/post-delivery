package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rjar2020/post-delivery/controller"
	"github.com/rjar2020/post-delivery/env"

	"github.com/rjar2020/post-delivery/consumers"
	"github.com/rjar2020/post-delivery/producers"
)

const example = `{
	"endpoint":{
	"method":"GET", 
	"url":"http://sample_domain_endpoint.com/data?icon={mascot}&image={coordinates}&foo={bar}"
	}, 
	"data":[
	{
	"icon":"Gopher", "coordinates":"https://blog.golang.org/gopher/gopher.png"
	} ]
	}`

func main() {
	log.Printf("####### Strating post delivery application #######")
	env.LoadEnv()
	controller.RegisterControllers()
	var wg sync.WaitGroup
	wg.Add(1)
	httpPort := ":" + os.Getenv(env.APIPort)
	topic := os.Getenv(env.KafkaPostBackTopic)
	producers.Produce(example, topic)
	go consumers.StartConsumer(topic, "postback-processor", &wg)
	producers.Produce(example, topic)
	http.ListenAndServe(httpPort, nil)
	wg.Wait()
}
