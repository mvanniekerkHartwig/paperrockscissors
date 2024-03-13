test:
	go test ./... -v

run:
	go run main.go

build:
	go build -o bin/main main.go

compile:
	rm -rf bin/*
	GOOS=darwin GOARCH=amd64 go build -o bin/rps-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/rps-darwin-arm64 main.go

	GOOS=linux GOARCH=amd64 go build -o bin/rps-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/rps-linux-arm64 main.go

	GOOS=windows GOARCH=amd64 go build -o bin/rps-windows-amd64 main.go
	GOOS=windows GOARCH=arm64 go build -o bin/rps-windows-arm64 main.go
