# Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o smr-api ./cmd/main.go

# Final (distroless) — alternativa: FROM scratch
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/smr-api /app/smr-api
# Se precisa de config por bind mount, garanta o path /app/config.json no compose
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app/smr-api"]
