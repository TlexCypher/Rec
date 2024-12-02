FROM golang:1.23.3

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
