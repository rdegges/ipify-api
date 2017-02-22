FROM golang:1.6
WORKDIR /var/www
ADD . /var/www

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/rdegges/ipify-api/api
RUN go get github.com/rs/cors
RUN go build -o ./server .

CMD ["./server"]
