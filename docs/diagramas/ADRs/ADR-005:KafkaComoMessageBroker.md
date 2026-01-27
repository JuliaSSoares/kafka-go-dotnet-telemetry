### ADR-005: Kafka como Message Broker

**Status:** ✅ Aprovado  
**Data:** 2026-01-20  
**Contexto:** Necessidade de streaming de eventos com alta vazão e durabilidade.

**Decisão:** Utilizar Apache Kafka como broker de mensagens.

**Consequências:**
- ✅ **Positivas:**
  - Throughput de 100k+ msg/s por partition
  - Retenção configurável (replay de eventos)
  - Particionamento para escalabilidade horizontal
  - Consumer Groups para load balancing
  - Garantia de ordem dentro de partições
  - Exatamente uma vez (exactly-once semantics) disponível
  - Ecosystem rico (Connect, Streams, KSQL)

- ❌ **Negativas:**
  - Complexidade operacional (Zookeeper até versão 3.x)
  - Overhead de infraestrutura (mínimo 3 brokers em produção)
  - Curva de aprendizado significativa
  - Não é adequado para mensagens com prioridades

**Alternativas Consideradas:**
1. **RabbitMQ** - Menor throughput, melhor para RPC patterns
2. **AWS SQS/SNS** - Vendor lock-in, custo por mensagem
3. **Redis Streams** - Menos maduro, sem ecosystem
