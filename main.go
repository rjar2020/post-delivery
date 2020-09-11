package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rjar2020/post-delivery/consumer"
	"github.com/rjar2020/post-delivery/controller"
	"github.com/rjar2020/post-delivery/env"
	"github.com/rjar2020/post-delivery/producer"
)

const example = `{
	"endpoint":{
	"method":"GET", 
	"url":"http://sample_domain_endpoint.com/data?tittle={icon}&image={coordinates}&foo={bar}"
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
	producer.Produce(example, topic)
	go consumer.StartPostBackConsumer(topic, "postback-processor", &wg)
	producer.Produce(example, topic)
	http.ListenAndServe(httpPort, nil)
	wg.Wait()
}
