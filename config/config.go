package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config struct holds the configuration settings.
type Config struct {
	HTTPPort string

	// Kafka Configuration
	KafkaBrokers              []string
	KafkaBrokersTest          []string
	KafkaMilestoneTopic       string
	KafkaMemoryTopic          string
	KafkaCommentTopic         string
	KafkaCustomEventTopic     string
	KafkaHistoricalEventTopic string
	KafkaMediaTopic           string

	LOG_PATH        string
	TimelineSvcAddr string
	MemorySvcAddr   string
}

// Load loads the configuration from environment variables.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":9090"))

	config.KafkaBrokers = cast.ToStringSlice(coalesce("KAFKA_BROKERS", []string{"kafka:9092"}))
	config.KafkaBrokersTest = cast.ToStringSlice(coalesce("KAFKA_BROKERS_Test", []string{"localhost:9092"}))

	// Kafka Topics
	config.KafkaMilestoneTopic = cast.ToString(coalesce("KAFKA_MILESTONE_TOPIC", "milestone_topic"))
	config.KafkaMemoryTopic = cast.ToString(coalesce("KAFKA_MEMORY_TOPIC", "memory_topic"))
	config.KafkaCommentTopic = cast.ToString(coalesce("KAFKA_COMMENT_TOPIC", "comment_topic"))
	config.KafkaCustomEventTopic = cast.ToString(coalesce("KAFKA_CUSTOM_EVENT_TOPIC", "custom_event_topic"))
	config.KafkaHistoricalEventTopic = cast.ToString(coalesce("KAFKA_HISTORICAL_EVENT_TOPIC", "historical_event_topic"))
	config.KafkaMediaTopic = cast.ToString(coalesce("KAFKA_MEDIA_TOPIC", "media_topic"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	config.TimelineSvcAddr = cast.ToString(coalesce("TIME_LINE_SERVICE_port", "timeline:9091"))
	config.MemorySvcAddr = cast.ToString(coalesce("MEMORY_SERVICE_port", "memory:9090"))
	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
