package domain

type Telemetria struct {
	ID        string  `json:"entregador_id"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Timestamp int64   `json:"timestamp"`
}
