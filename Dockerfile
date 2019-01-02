# build image

FROM golang:alpine AS build

WORKDIR /go/src/github.com/rdegges/ipify-api
COPY . /go/src/github.com/rdegges/ipify-api
RUN go build

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/go/src/github.com/rdegges/ipify-api/ipify-api"]

# live image

FROM build AS live
EXPOSE 3000/tcp

RUN apk upgrade \
   && apk add --update dumb-init \
   && rm /var/cache/apk/*

COPY --from=build /go/src/github.com/rdegges/ipify-api/ipify-api /go/bin/

CMD ["/go/bin/ipify-api"]
