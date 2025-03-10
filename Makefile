CC = go
OUTPUT = bin/app
INPUT = cmd/server/main.go

build:
	$(CC) build -o $(OUTPUT) $(INPUT)
run:
	$(MAKE) build
	$(OUTPUT)
deps:
	go mod 