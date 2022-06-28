package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer(configMap *ckafka.ConfigMap) (*ckafka.Producer, error) {
	var newProducer, err = ckafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}

	return newProducer, nil
}

func Publish(producer *ckafka.Producer, msg, topic string) error {
	var kafkaMessage = &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	var err = producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}

	return nil
}
