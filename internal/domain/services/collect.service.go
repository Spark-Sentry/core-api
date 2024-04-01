package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/infrastructure/kafka"
	"encoding/json"
	"fmt"
)

// CollectService is responsible for processing collected data based on parameters.
type CollectService struct {
	kafkaProducer kafka.Producer // Assumed to be your adapted Kafka producer
}

// NewCollectService creates a new instance of CollectService.
func NewCollectService(kafkaProducer kafka.Producer) *CollectService {
	return &CollectService{kafkaProducer: kafkaProducer}
}

// ProcessParameterValuesBatch processes the batch of parameter values and sends it to Kafka.
func (s *CollectService) ProcessParameterValuesBatch(dataBatch dto.ParameterValuesBatch) error {
	for _, value := range dataBatch.Values {
		messageData := struct {
			ParameterID string  `json:"parameterId"`
			Timestamp   string  `json:"timestamp"`
			Value       float64 `json:"value"`
		}{
			ParameterID: dataBatch.ParameterID,
			Timestamp:   value.Timestamp,
			Value:       value.Value,
		}

		message, err := json.Marshal(messageData)
		if err != nil {
			fmt.Printf("Error marshalling the message data: %v\n", err)
			continue // or return the error, depending on your error handling strategy
		}

		if err := s.kafkaProducer.SendMessage("parameter-values-topic", message); err != nil {
			fmt.Printf("Failed to send message to Kafka: %v\n", err)
			return err
		}
	}

	return nil
}
