FROM golang:1.23.6 AS builder

WORKDIR /online_market_project

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o myapp ./cmd/main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /online_market_project/myapp .

COPY --from=builder /online_market_project/cmd/config.yaml .

EXPOSE 8080

CMD ["./myapp"]