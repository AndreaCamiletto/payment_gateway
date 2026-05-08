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
RUN adduser -D -u 10001 appuser
USER 10001

WORKDIR /home/appuser/

# proprietario del file e\' l'utente non-root
COPY --from=builder --chown=10001:10001 /payment-gateway .

# Eseguo come utente non-root
USER 10001

EXPOSE 8080

CMD ["./payment-gateway"]