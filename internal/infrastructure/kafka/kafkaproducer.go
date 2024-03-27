// internal/infrastructure/kafka/kafkaproducer.go

package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// Producer KafkaProducer is an implementation of KafkaProducerInterface using confluent-kafka-go.
type Producer struct {
	producer *kafka.Producer
}

// NewKafkaProducer creates a new instance of KafkaProducer.
func NewKafkaProducer(brokerList string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerList})
	if err != nil {
		return nil, err
	}
	log.Println("✅ successfully connected to kafka")

	return &Producer{producer: p}, nil
}

// SendMessage sends a message to the specified Kafka topic.
func (kp *Producer) SendMessage(topic string, message []byte) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
}
