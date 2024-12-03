FROM golang:1.23.3

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

RUN apt-get update && apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
