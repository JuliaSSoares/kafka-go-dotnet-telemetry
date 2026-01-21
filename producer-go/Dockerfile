# Usamos uma imagem que já tem o Go e o compilador C instalados
FROM golang:1.21-bullseye

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código
COPY . .

# Compila e roda
CMD ["go", "run", "main.go"]