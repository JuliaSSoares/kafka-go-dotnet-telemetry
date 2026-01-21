# 📍 kafka-go-dotnet-telemetry

Sistema de alta performance para processamento de eventos de telemetria (GPS) em tempo real. Este projeto demonstra uma arquitetura orientada a eventos (EDA) utilizando **Apache Kafka** para desacoplar a ingestão de dados da lógica de negócio.

## 🏗️ Arquitetura do Sistema

O projeto simula um ecossistema de logística onde milhares de dispositivos enviam coordenadas geográficas simultaneamente:

1.  **Producer (Go):** Um serviço de alta performance responsável por simular e enviar eventos de telemetria para o Kafka. Escolhido pela eficiência do Go em lidar com concorrência e baixa latência.
2.  **Streaming Platform (Kafka):** Atua como o "cérebro" do sistema, garantindo a persistência, ordenação e distribuição dos eventos.
3.  **Consumer (.NET 8):** Um serviço robusto que processa os dados recebidos, simulando regras de negócio como alertas de proximidade ou cálculos de rota.

## 🛠️ Tecnologias Utilizadas

| Componente | Tecnologia | Papel no Projeto |
| :--- | :--- | :--- |
| **Linguagem (Producer)** | Go (Golang) | Ingestão massiva de dados |
| **Linguagem (Consumer)** | .NET 8 (C#) | Processamento e regras de negócio |
| **Message Broker** | Apache Kafka | Event Streaming e Mensageria |
| **Containerização** | Docker & Compose | Orquestração da infraestrutura local |

## 🚀 Como Executar

### Pré-requisitos
* Docker e Docker Compose instalados.
* Go 1.21+ (para rodar o producer localmente).
* .NET SDK 8 (para rodar o consumer localmente).

### 1. Subir a Infraestrutura
Na raiz do projeto, execute o comando para subir o Kafka e o Zookeeper:
```bash
docker-compose up -d
