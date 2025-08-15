BIN_NAME=terraform-provider-slack
PKG=github.com/ryanlwinters/terraform-provider-slack
VERSION?=dev

.PHONY: build test lint docs acc clean

build:
	go build -ldflags "-X $(PKG)/provider.Version=$(VERSION)" -o $(BIN_NAME)

test:
	go test ./...

docs:
	@echo "Docs generation not yet configured"

acc:
	TF_ACC=1 go test ./... -v -timeout 30m

clean:
	rm -f $(BIN_NAME)


