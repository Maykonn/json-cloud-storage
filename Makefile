run: build
	./bin/build

build: clear fmt
	go build -o ./bin/build ./main.go

dependencies:
	go get -u github.com/aws/aws-sdk-go

clear:
	-rm ./build

fmt:
	gofmt -w .

.PHONY: run build dependencies clear fmt