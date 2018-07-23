GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin

all: ./bin/influx_proxy

./bin/influx_proxy: 
	go build -o ./bin/influx_proxy ./cmd/influx_proxy

./vendor:
	test -d ./vendor/src || mkdir -p ./vendor/src
	glide install
	test -d ./vendor/src || (mkdir ./src && mv ./vendor/* ./src && mv ./src ./vendor)

clean:
	rm -rf bin

.PHONY: clean update
