build:
	go build -o ./bin/blockgo

run: build
	./bin/blockgo

test:
	go test -v ./...