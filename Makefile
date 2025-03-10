all: gopass

bin:
	@echo "Creating bin directory..."
	@mkdir -p bin

gopass: bin
	@echo "Creating binary file gopath..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/gopass ./cmd/gopass
	@echo "gopass created into bin directory"

check-upx:
	@command -v upx > /dev/null 2>&1 || { \
		echo "upx is not installed. Please install it."; \
		exit 1; \
	}

install: check-upx gopass
	@echo "Installing gopass to $(GOBIN)"
	@upx --best ./bin/gopass
	@cp ./bin/gopass $(GOBIN)
	@echo "gopass installed as $(GOBIN)/gopass"

clean:
	@echo "Cleaning..."
	@rm -rf bin
