FROM golang:1.18.2 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd 

EXPOSE 8000

CMD go run cmd/main.go