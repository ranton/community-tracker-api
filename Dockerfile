FROM golang:1.18.2 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o community-tracker-api ./cmd/main.go
EXPOSE 8000

CMD ["./community-tracker-api"]
