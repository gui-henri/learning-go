FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

EXPOSE 8090

CMD ["go", "run", "main.go"]