FROM golang:1.8 as builder

WORKDIR /app

COPY vendor/github.com /usr/local/go/src/github.com
COPY api /usr/local/go/src/github.com/DNSFilter/ipify-api/api
COPY models /usr/local/go/src/github.com/DNSFilter/ipify-api/models
COPY . .
RUN go build -o ip && chmod +x /app/ip

CMD ["/app/ip"]