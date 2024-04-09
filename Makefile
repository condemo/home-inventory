binary-name=home-inventory

build:
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/main.go
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-windows.exe ./cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/main.go

run: build
	@./bin/${binary-name}-linux

clean:
	@rm -rf ./bin/*
	@go clean
