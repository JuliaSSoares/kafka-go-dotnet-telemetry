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

// PublishRaw é o que o RELAY vai usar. Ele recebe o []byte vindo do Outbox.
func (kp *KafkaProducer) PublishRaw(topic string, payload []byte) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
}

// Close fecha o producer (importante para não deixar conexões penduradas)
func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}
