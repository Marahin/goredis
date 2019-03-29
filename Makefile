build: install-dependencies
	mkdir -p bin/
	go build -o bin/goredis
	chmod +x bin/goredis

run: build
	bin/goredis

install-dependencies:
	dep ensure

test: install-dependencies
	go test ./...
