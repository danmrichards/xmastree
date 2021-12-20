GOARCH=amd64

.PHONY: build
build: linux windows darwin

.PHONY: linux
linux:
	GOOS=linux go build -ldflags="-s -w" -o bin/xmastree-linux-${GOARCH} ./cmd/main.go

.PHONY: windows
windows:
	GOOS=windows go build -ldflags="-s -w" -o bin/xmastree-windows-${GOARCH}.exe ./cmd/main.go

.PHONY: darwin
darwin:
	GOOS=darwin go build -ldflags="-s -w" -o bin/xmastree-darwin-${GOARCH} ./cmd/main.go