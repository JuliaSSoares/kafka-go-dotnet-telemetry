package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Telemetria struct {
	ID        string  `json:"entregador_id"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
	}

	p, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Erro ao criar o produtor: %s\n", err)
		return
	}
	defer p.Close()

	topico := "telemetria-entregadores"

	fmt.Println("🚀 Produtor Go iniciado! Enviando dados...")

	for {
		// Criando um dado fake de GPS
		msg := Telemetria{
			ID:        "entregador-01",
			Lat:       -23.5505,
			Long:      -46.6333,
			Timestamp: time.Now().Unix(),
		}

		payload, _ := json.Marshal(msg)

		// Envia para o tópico
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topico, Partition: kafka.PartitionAny},
			Value:          payload,
		}, nil)

		if err != nil {
			fmt.Printf("Erro ao enviar: %s\n", err)
		} else {
			fmt.Printf("✅ Mensagem enviada: %s\n", string(payload))
		}

		time.Sleep(2 * time.Second) // Espera 2 segundos para não inundar o Kafka
	}
}
