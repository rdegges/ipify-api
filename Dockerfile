FROM golang:1.9

RUN mkdir -p /go/src/github.com/rdegges/ipify-api

COPY main.go /go/src/github.com/rdegges/ipify-api
COPY api /go/src/github.com/rdegges/ipify-api/api
COPY models /go/src/github.com/rdegges/ipify-api/models

WORKDIR /go/src/github.com/rdegges/ipify-api

RUN go get -v -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /

COPY --from=0 /go/src/github.com/rdegges/ipify-api/app .

CMD ["./app"]
