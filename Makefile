.PHONY: tidy build run build_and_run run_local clean

tidy:
	go mod tidy

build:
	@mkdir -p bin
	go build -o bin/rapidcache cmd/rapidcache/main.go
	@echo "-->Built server binary at bin/rapidcache"

run:
	@echo "-->Running binary ..."
	@chmod +x bin/rapidcache
	./bin/rapidcache

build_and_run: build run

run_local:
	go run cmd/rapidcache/main.go

clean:
	rm -rf bin/
	@echo "-->Deleted binary"
