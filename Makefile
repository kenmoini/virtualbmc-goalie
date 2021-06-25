test:
	go build -v -o dist/vbmcgoalie main.go

build:
	go build -v -o dist/vbmcgoalie main.go

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o dist/vbmcgoalie-linux-amd64 main.go

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -v -o dist/vbmcgoalie-dawrwin-amd64 main.go

run:
	go run main.go
