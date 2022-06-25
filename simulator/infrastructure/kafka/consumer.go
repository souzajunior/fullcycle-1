package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"simulator-fc1/config"
)

// KafkaConsumer represents the kafka consumer
type KafkaConsumer struct {
	Msg chan *ckafka.Message
}

func newKafkaConsumer(configMap *ckafka.ConfigMap) (kafkaConsumer *ckafka.Consumer, err error) {
	kafkaConsumer, err = ckafka.NewConsumer(configMap)
	if err != nil {
		return
	}

	return
}

// Consume is responsible to consume the data from kafka topic
func (k *KafkaConsumer) Consume() {
	var configKafkaMap = ckafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().KafkaBootstrapServers,
		"group.id":          config.GetConfig().KafkaConsumerGroupID,
	}

	var kafkaConsumer, err = newKafkaConsumer(&configKafkaMap)
	if err != nil {
		log.Fatalf("error initializing a new consumer: %s", err.Error())
	}

	err = kafkaConsumer.Subscribe(config.GetConfig().KafkaReadTopic, nil)
	if err != nil {
		log.Fatalf("error subscribing into read topic: %s", err.Error())
	}

	log.Println("Kafka consumer has been started")

	for {
		var kfkMsg, err = kafkaConsumer.ReadMessage(-1)
		if err != nil {
			log.Printf("error while reading message: %s", err.Error())
		}

		k.Msg <- kfkMsg
	}

}
