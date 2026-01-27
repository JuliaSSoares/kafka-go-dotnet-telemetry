### ADR-002: Escolha de Go para o Producer

**Status:** ✅ Aprovado  
**Data:** 2026-01-20  
**Contexto:** Necessidade de alta concorrência para simular milhares de sensores GPS simultâneos.

**Decisão:** Utilizar Go (Golang) para o Producer (Simulator + Relay).

**Consequências:**
- ✅ **Positivas:**
  - Goroutines para concorrência leve (10.000+ sensores simultâneos)
  - Performance nativa sem overhead de runtime
  - Baixo consumo de memória (~25MB por instância)
  - Excelente suporte a I/O não-bloqueante
  - Compilação estática (binário único)

- ❌ **Negativas:**
  - Curva de aprendizado para desenvolvedores sem experiência em Go
  - Ecossistema menor comparado a Java/C#
  - Menos ferramentas de debugging visual

**Alternativas Consideradas:**
1. **Java com Virtual Threads** - Mais verboso, maior consumo de memória
2. **Node.js** - Single-threaded, menor throughput
3. **Rust** - Curva de aprendizado muito íngreme
