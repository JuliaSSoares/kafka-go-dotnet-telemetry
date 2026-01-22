-- Tabela de Histórico (Source of Truth)
CREATE TABLE IF NOT EXISTS telemetrias (
    id SERIAL PRIMARY KEY,
    entregador_id VARCHAR(50) NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    long DOUBLE PRECISION NOT NULL,
    timestamp BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de Outbox (Fila para o Relay)
CREATE TABLE IF NOT EXISTS outbox_messages (
    id SERIAL PRIMARY KEY,
    payload JSONB NOT NULL,
    topic VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index para o Relay não ficar lento conforme a tabela cresce
CREATE INDEX idx_outbox_status ON outbox_messages(status) WHERE status = 'PENDING';