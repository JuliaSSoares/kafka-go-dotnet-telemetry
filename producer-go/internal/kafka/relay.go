package kafka

import (
	"database/sql"
	"time"
)

func StartOutboxRelay(db *sql.DB, producer *KafkaProducer) {
	for {
		rows, _ := db.Query("SELECT id, payload FROM outbox_messages WHERE status = 'PENDING' LIMIT 10")

		for rows.Next() {
			var id int
			var payload []byte
			var topic string
			rows.Scan(&id, &payload, &topic)

			err := producer.PublishRaw(topic, payload)

			if err == nil {
				db.Exec("UPDATE outbox_messages SET status = 'PROCESSED' WHERE id = $1", id)
			}
		}
		rows.Close()
		time.Sleep(2 * time.Second) // "Consumindo aos poucos"
	}
}
