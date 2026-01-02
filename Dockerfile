# Estágio de Build
FROM golang:alpine AS builder

WORKDIR /app

# Cache das dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte e compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Estágio Final (Imagem leve)
FROM alpine:latest

WORKDIR /root/

# Copia o binário do estágio anterior
COPY --from=builder /app/main .

# Porta que a aplicação vai rodar (ajuste se necessário)
EXPOSE 8080

CMD ["./main"]