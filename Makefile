all: gopass

bin:
	@echo "Creating bin directory..."
	@mkdir -p bin

gopass: bin
	@echo "Creating binary file gopath..."
	@go build -o ./bin/gopass ./cmd/gopass
	@echo "gopass created into bin directory"

clean:
	@echo "Cleaning..."
	@rm -rf bin
