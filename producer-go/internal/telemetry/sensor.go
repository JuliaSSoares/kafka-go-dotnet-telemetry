package telemetry

import (
	"math/rand"
	"time"

	"github.com/julia.soares/producer-go/internal/domain"
)

// Simulador de GPS
type GPSSensor struct {
	EntregadorID string
	CurrentLat   float64
	CurrentLong  float64
}

func NewGPSSensor(id string) *GPSSensor {
	return &GPSSensor{
		EntregadorID: id,
		CurrentLat:   -23.5505, // São Paulo
		CurrentLong:  -46.6333,
	}
}

// Read gera um novo ponto de telemetria simulando movimento
func (s *GPSSensor) Read() domain.Telemetria {
	// Simula um pequeno deslocamento aleatório
	s.CurrentLat += (rand.Float64() - 0.5) * 0.001
	s.CurrentLong += (rand.Float64() - 0.5) * 0.001

	return domain.Telemetria{
		ID:        s.EntregadorID,
		Lat:       s.CurrentLat,
		Long:      s.CurrentLong,
		Timestamp: time.Now().Unix(),
	}
}
