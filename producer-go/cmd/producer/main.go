package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julia.soares/producer-go/internal/kafka"
	"github.com/julia.soares/producer-go/internal/repository"
	"github.com/julia.soares/producer-go/internal/telemetry"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DB_SOURCE")
	if connStr == "" {
		connStr = "postgresql://user_geo:password_geo@localhost:5432/geoprocessing?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro DB:", err)
	}
	defer db.Close()

	kafkaServer := os.Getenv("KAFKA_SERVER")
	if kafkaServer == "" {
		kafkaServer = "localhost:9092"
	}

	producer, err := kafka.NewKafkaProducer(kafkaServer)
	if err != nil {
		log.Fatal("Erro Kafka:", err)
	}
	defer producer.Close()

	repo := repository.NewRepository(db)
	sensor := telemetry.NewGPSSensor("entregador-01")

	go func() {
		for {
			dado := sensor.Read()
			err := repo.SaveTelemetriaWithOutbox(context.Background(), dado)
			if err != nil {
				log.Println("Erro ao processar sensor:", err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	go kafka.StartOutboxRelay(db, producer)

	log.Println("🚀 Sistema rodando. Sensor capturando e Relay monitorando...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Encerrando aplicação...")
}
