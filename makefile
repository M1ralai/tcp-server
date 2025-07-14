compile: build run

build:
	go build -o bin/output  cmd/main.go 

run:
	./bin/output