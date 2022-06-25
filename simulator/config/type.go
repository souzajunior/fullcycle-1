package config

// Config represents the config data for the service
type Config struct {
	KafkaBootstrapServers string
	KafkaConsumerGroupID  string
	KafkaReadTopic        string
}

const (
	envkafkaBootstrapServers = "KAFKA_BOOTSTRAP_SERVERS"
	envkafkaConsumerGroupID  = "KAFKA_CONSUMER_GROUP_ID"
	envkafkaReadTopic        = "KAFKA_READ_TOPIC"
)
