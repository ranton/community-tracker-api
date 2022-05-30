FROM golang:1.18.2 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o community-tracker-api ./cmd/main.go


FROM alpine as build 
WORKDIR /app
COPY --from=development /app/.env ./
COPY --from=development /app/community-tracker-api ./
EXPOSE 8000

CMD ["./community-tracker-api"]
