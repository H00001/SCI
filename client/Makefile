LIN_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
MAC_ENV=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
BUILD=go build

build-linux:
	$(LIN_ENV) $(BUILD)

build-mac:
	$(MAC_ENV) $(BUILD)
