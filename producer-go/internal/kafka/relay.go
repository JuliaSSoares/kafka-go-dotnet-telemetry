package kafka

import (
	"database/sql"
	"time"
)

func StartOutboxRelay(db *sql.DB, producer *KafkaProducer) {
	for {
		rows, err := db.Query("SELECT id, payload, topic FROM outbox_messages LIMIT 10")
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		for rows.Next() {
			var id int
			var payload []byte
			var topic string

			if err := rows.Scan(&id, &payload, &topic); err != nil {
				continue
			}

			if topic == "" {
				topic = "telemetria-topic"
			}

			err := producer.PublishRaw(topic, payload)

			if err == nil {
				db.Exec("DELETE FROM outbox_messages WHERE id = $1", id)
			}
		}
		rows.Close()
		time.Sleep(2 * time.Second)
	}
}
