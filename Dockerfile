FROM golang:1.6
WORKDIR /var/www
ADD . /var/www

RUN go get github.com/julienschmidt/httprouter github.com/rdegges/ipify-api/api github.com/rs/cors
RUN go build -o ./server .

CMD ["./server"]
