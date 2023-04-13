run:
	go run main.go

mod:
	go mod tidy
	go mod vendor

build_darwin:
	@echo "Building for macOS..."
	@GOOS=darwin go build -o bin/$(v)/macOS main.go
build_linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/$(v)/linus main.go
build_windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/$(v)/windows.exe main.go
build: build_darwin build_linux build_windows