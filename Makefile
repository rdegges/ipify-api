TAG=latest
BIN=ipify-api
IMAGE=unchartedsky/$(BIN)

build:
	dep ensure
	go build -o bin/$(BIN) .

image: build
	dep ensure
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN) .
	docker build -t $(IMAGE):$(TAG) .

deploy: image
	docker push $(IMAGE):$(TAG)

.PHONY: clean

clean:
	rm -rf bin/

cleanall: clean
	rm -rf vendor/

