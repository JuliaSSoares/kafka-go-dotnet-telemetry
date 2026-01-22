package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/julia.soares/producer-go/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveTelemetriaWithOutbox(ctx context.Context, t domain.Telemetria) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryHistory := `INSERT INTO telemetrias (entregador_id, lat, long, timestamp) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, queryHistory, t.ID, t.Lat, t.Long, t.Timestamp)
	if err != nil {
		return err
	}

	payload, _ := json.Marshal(t)

	queryOutbox := `INSERT INTO outbox_messages (payload, topic) VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, queryOutbox, payload, "telemetria-topic")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func SaveToOutbox(db *sql.DB, t domain.Telemetria) error {
	payload, _ := json.Marshal(t)

	_, err := db.Exec(
		"INSERT INTO outbox_messages (payload, status) VALUES ($1, 'PENDING')",
		payload,
	)
	return err
}
