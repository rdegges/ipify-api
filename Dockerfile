FROM golang:1.4.2

WORKDIR /gopath/src/github.com/rdegges/ipify-api
ADD . /gopath/src/github.com/rdegges/ipify-api/

RUN go get github.com/tools/godep
RUN godep restore
RUN godep go install

EXPOSE 3000

CMD []
ENTRYPOINT ["/gopath/bin/ipify-api"]
