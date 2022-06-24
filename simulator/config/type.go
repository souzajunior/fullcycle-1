package config

// Config represents the config data for the service
type Config struct {
	KafkaBootstrapServers string
	KafkaConsumerGroupID  string
}

const (
	envkafkaBootstrapServers = "KAFKA_BOOTSTRAP_SERVERS"
	envkafkaConsumerGroupID  = "KAFKA_CONSUMER_GROUP_ID"
)
