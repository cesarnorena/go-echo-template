FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-echo-template ./cmd/.

EXPOSE 8080

CMD ["/app/go-echo-template"]
