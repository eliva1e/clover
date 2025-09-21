all: build-cli build-image

build-cli:
	GOOS=windows go build -o dist-cli/CloverCLI_windows.exe cmd/cli/main.go
	GOOS=darwin go build -o dist-cli/CloverCLI_macos cmd/cli/main.go
	GOOS=linux go build -o dist-cli/CloverCLI_linux cmd/cli/main.go

build-image:
	docker build -t eliva1e/clover --platform linux/amd64,linux/arm64 .
