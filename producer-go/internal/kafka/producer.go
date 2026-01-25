package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func (kp *KafkaProducer) PublishRaw(topic string, payload []byte) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
}

func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}
