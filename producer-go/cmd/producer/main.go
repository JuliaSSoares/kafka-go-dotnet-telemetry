package main

import (
	"log"
	"time"

	"github.com/julia.soares/producer-go/internal/domain"
	"github.com/julia.soares/producer-go/internal/kafka"
)

func main() {
	// Inicializa o produtor
	client, err := kafka.NewKafkaProducer("kafka:29092")
	if err != nil {
		log.Fatal("Falha ao iniciar Kafka:", err)
	}

	for {
		dado := domain.Telemetria{
			ID:  "entregador-01",
			Lat: -23.5, Long: -46.6,
			Timestamp: time.Now().Unix(),
		}

		if err := client.Publish("telemetria-entregadores", dado); err != nil {
			log.Println("Erro ao publicar:", err)
		} else {
			log.Printf("📡 Enviado: %s", dado.ID)
		}

		time.Sleep(2 * time.Second)
	}
}
