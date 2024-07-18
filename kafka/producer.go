package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/time_capsule/Api-Gateway-Timecapsule/config"
)

// Producer produces Kafka messages.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer creates a new Producer instance for the given topic.
func NewProducer(topic string) *Producer {
	cfg := config.Load()
	CreateTopic(cfg.KafkaBrokersTest, topic)
	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.KafkaBrokers...),
		Topic:        topic,
		RequiredAcks: kafka.RequireOne,
		Balancer:     &kafka.LeastBytes{},
	}
	return &Producer{writer: writer}
}

// ProduceMessage produces a message with the given key and value.
func (p *Producer) ProduceMessage(ctx context.Context, key string, message interface{}) error {
	value, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("failed to write message to Kafka: %w", err)
	}

	return nil
}

// Helper functions to create, delete, and produce messages to a Kafka topic
func CreateTopic(brokers []string, topic string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	return nil
}
