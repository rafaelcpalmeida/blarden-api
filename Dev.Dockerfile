FROM golang:1.13.1

COPY . /app

WORKDIR /app

RUN trap $(go mod init blarden-api)

RUN go build -o bin/api cmd/main.go

CMD ["go", "run", "cmd/main.go"]
