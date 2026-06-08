.PHONY: all test bench cover fmt vet clean

all: fmt vet test

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f coverage.out
