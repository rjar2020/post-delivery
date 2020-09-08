package env

import (
	"log"
	"os"

	"github.com/magiconair/properties"
)

const (
	//KafkaBrokers define kafka brokers URL
	KafkaBrokers = "KAFKA_BROKERS"
	//KafkaConsumerOffsetReset define auto.offset.reset kafka consumer config/policy
	KafkaConsumerOffsetReset = "KAFKA_CONSUMER_OFFSET_RESET"
	//KafkaProducersAck define acks policy/config for kafka producers
	KafkaProducersAck = "KAFKA_PRODUCER_ACKS"
	//APIPort is the port where the http server will be started
	APIPort = "API_PORT"
	//KafkaPostBackTopic is the topic where the postback messages are delivered
	KafkaPostBackTopic = "KAFKA_POSTBACK_TOPIC"
)

//LoadEnv retrieve environment variables
func LoadEnv() {
	p := properties.MustLoadFile("./env.properties", properties.UTF8)
	log.Printf("Properties: %s", p)
	os.Setenv(KafkaBrokers, p.MustGetString(KafkaBrokers))
	os.Setenv(KafkaConsumerOffsetReset, p.MustGetString(KafkaConsumerOffsetReset))
	os.Setenv(KafkaProducersAck, p.MustGetString(KafkaProducersAck))
	os.Setenv(APIPort, p.MustGetString(APIPort))
	os.Setenv(KafkaPostBackTopic, p.MustGetString(KafkaPostBackTopic))
}
