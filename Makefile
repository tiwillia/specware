.PHONY: build clean test

build:
	go build -o specware .

clean:
	rm -f specware

test:
	go test ./...
	go test ./tests/...