package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/julia.soares/producer-go/internal/domain"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(server string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: p}, nil
}

func (kp *KafkaProducer) Publish(topic string, data domain.Telemetria) error {
	payload, _ := json.Marshal(data)
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
}
