VERBOSE=1

run: build
	./bin/build

build: clear fmt
	go build -o ./bin/build ./main.go

test:
	go test ./...

dependencies:
	go get -u github.com/aws/aws-sdk-go

clear:
	-rm ./build

fmt:
	gofmt -w .

.PHONY: run build test dependencies clear fmt