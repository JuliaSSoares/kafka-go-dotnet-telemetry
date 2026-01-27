
### ADR-001: Uso do Transactional Outbox Pattern

**Status:** ✅ Aprovado  
**Data:** 2026-01-20  
**Contexto:** Necessidade de garantir entrega de eventos sem perda de dados em caso de falha do Kafka.

**Decisão:** Implementar Outbox Pattern com PostgreSQL como buffer transacional.

**Consequências:**
- ✅ **Positivas:**
  - Atomicidade garantida entre escrita no banco e publicação no Kafka
  - Resiliência a falhas temporárias do Kafka
  - Possibilidade de retry automático
  - Auditoria completa de eventos

- ❌ **Negativas:**
  - Latência adicional (polling interval ~1s)
  - Necessidade de processo separado (Relay)
  - Maior uso de disco no PostgreSQL

**Alternativas Consideradas:**
1. **Publicação direta no Kafka** - Descartado por falta de atomicidade
2. **CDC (Change Data Capture)** - Complexidade operacional elevada
3. **Dual Write** - Risco de inconsistência

