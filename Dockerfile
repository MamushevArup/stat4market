FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/config ./config
COPY --from=builder /app/.env ./
COPY --from=builder /app/migration ./migration

CMD ["./main"]