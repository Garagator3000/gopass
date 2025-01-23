all: gopass

bin:
	mkdir -p bin

gopass: bin
	go build -o ./bin/gopass ./cmd/gopass
