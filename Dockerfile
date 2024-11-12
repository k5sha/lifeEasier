FROM golang:alpine AS builder

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go
EXPOSE 8080

CMD ["sh", "-c", "goose -dir ./internal/storage/migrations postgres \"postgresql://postgres:postgres@db:5432/life_easier_db?sslmode=disable\" up && /main"]