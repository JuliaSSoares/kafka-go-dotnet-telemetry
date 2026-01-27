### ADR-003: Escolha de .NET para o Consumer

**Status:** ✅ Aprovado  
**Data:** 2026-01-20  
**Contexto:** Necessidade de processamento robusto com suporte a ORM e padrões empresariais.

**Decisão:** Utilizar .NET 10 com Worker Service para o Consumer.

**Consequências:**
- ✅ **Positivas:**
  - Entity Framework Core para persistência robusta
  - MediatR para CQRS e desacoplamento
  - Suporte nativo a DI (Dependency Injection)
  - Ecossistema maduro para aplicações empresariais
  - Performance excelente (.NET 10 com AOT)
  - Strong typing e nullability check

- ❌ **Negativas:**
  - Maior uso de memória comparado a Go (~150MB)
  - Runtime maior (Docker image ~200MB)
  - Startup time ligeiramente maior

**Alternativas Consideradas:**
1. **Go também no Consumer** - Menos produtividade, sem ORM robusto
2. **Java com Spring** - Mais pesado e verboso
3. **Python** - Performance inferior, GIL limitations
