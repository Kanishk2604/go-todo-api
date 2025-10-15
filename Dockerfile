# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates tzdata
RUN addgroup -g 1001 -S appuser && adduser -u 1001 -S appuser -G appuser

WORKDIR /app
COPY --from=builder /app/main .
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]
