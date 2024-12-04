FROM golang:1.23.3 AS builder

WORKDIR /go/src


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# development
FROM golang:1.23.3 AS dev

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

RUN apt-get update && apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

COPY --from=builder /go/src/main .

COPY --from=builder /go/src/.air.toml .

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]

# production
FROM debian:bullseye-slim AS prod

WORKDIR /app

COPY --from=builder /go/src/main .

EXPOSE 8000

CMD ["./main"]

