# STAGE 1: Build
FROM golang:1.26-alpine AS builder

# Installiamo le dipendenze minime per compilare (se necessarie)
RUN apk add --no-cache git

WORKDIR /app

# Copio i file dei moduli
COPY go.mod ./


# Copio il resto del codice
COPY . .

# Compilazione dell'eseguibile statico
RUN CGO_ENABLED=0 GOOS=linux go build -o /payment-gateway ./cmd/server

FROM alpine:3.19

# utente non-root per sicurezza
RUN adduser -D -g '' appuser

WORKDIR /home/appuser/

COPY --from=builder /payment-gateway .

# proprietario del file e\' l'utente non-root
RUN chown appuser:appuser payment-gateway

# Eseguo come utente non-root
USER appuser

EXPOSE 8080

CMD ["./payment-gateway"]