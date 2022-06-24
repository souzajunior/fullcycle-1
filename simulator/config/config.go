package config

import (
	"log"
	"os"
)

var config *Config

// LoadConfig loads the config necessary data into config struct
func LoadConfig() {
	if config != nil {
		log.Println("Config has already been loaded")
		return
	}

	config = new(Config)

	var kafkaBootstrapServers = os.Getenv(envkafkaBootstrapServers)
	if kafkaBootstrapServers == "" {
		log.Fatalf("%s was not configured in environment variables", envkafkaBootstrapServers)
	}

	var kafkaConsumerGroupID = os.Getenv(envkafkaConsumerGroupID)
	if kafkaConsumerGroupID == "" {
		log.Fatalf("%s was not configured in environment variables", envkafkaConsumerGroupID)
	}

	config.KafkaBootstrapServers = kafkaBootstrapServers
	config.KafkaConsumerGroupID = kafkaConsumerGroupID
}

// GetConfig returns the config data
func GetConfig() *Config {
	if config == nil {
		log.Fatal("Config data was not loaded")
	}

	return config
}
