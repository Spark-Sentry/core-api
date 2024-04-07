package interfaces

// KafkaProducerInterface defines methods for sending messages to Kafka topics.
type KafkaProducerInterface interface {
	// SendMessage sends a message to a specific Kafka topic.
	SendMessage(topic string, message []byte) error
}
