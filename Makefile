CC = go
OUTPUT = bin/app
INPUT = cmd/server/main.go

build:
	$(CC) build -o $(OUTPUT) $(INPUT)
run:
	$(MAKE) build
	$(OUTPUT)
build-js:
	npm --prefix ./web run build:js

build-css:
	npm --prefix ./web run build:css
deps:
	go mod download