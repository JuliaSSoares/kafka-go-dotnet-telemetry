### ADR-004: Redis como Fast Layer

**Status:** ✅ Aprovado  
**Data:** 2026-01-20  
**Contexto:** Necessidade de acesso O(1) à última posição conhecida de cada motorista.

**Decisão:** Utilizar Redis como camada de cache para estado atual.

**Consequências:**
- ✅ **Positivas:**
  - Latência sub-milissegundo (< 1ms p99)
  - Estruturas de dados ricas (Hashes, Sets, Sorted Sets)
  - TTL automático para expiração (evita memory leaks)
  - Pub/Sub para eventos em tempo real
  - Suporte a Lua scripts para operações atômicas
  - Replicação master-slave disponível

- ❌ **Negativas:**
  - Dados voláteis (não é source of truth)
  - Necessidade de warm-up após restart
  - Custo de memória RAM (~1KB por chave)
  - Sem suporte a queries complexas

**Alternativas Consideradas:**
1. **PostgreSQL com índices** - Latência > 10ms, não escalável
2. **Memcached** - Menos features, sem TTL por chave
3. **In-Memory do próprio .NET** - Não compartilhado entre instâncias
