build: 
	go build -o bin/chainStore

run: build
	./bin/chainStore

.PHONY: run

test: 
	go test -v ./...

