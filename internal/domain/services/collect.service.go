package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/infrastructure/kafka"
	"encoding/json"
)

// CollectService is responsible for processing collected data.
type CollectService struct {
	kafkaProducer kafka.Producer
}

// NewCollectService creates a new instance of CollectService.
func NewCollectService(kafkaProducer kafka.Producer) *CollectService {
	return &CollectService{kafkaProducer: kafkaProducer}
}

// ProcessData processes the collected data and sends it to Kafka.
func (s *CollectService) ProcessData(data dto.CollectData) error {
	// Conversion de data en message []byte
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Envoi du message à un sujet Kafka
	return s.kafkaProducer.SendMessage("yourTopic", message)
}
