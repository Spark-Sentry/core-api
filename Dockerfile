FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sparksentry ./cmd/sparksentry/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/sparksentry .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./sparksentry"]
