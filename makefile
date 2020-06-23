default:
	go build -o bin/simple-server -v cmd/simple-server/main.go
	go build -o bin/data-creator -v cmd/data-creator/main.go

simple-server:
	go build -o bin/simple-server -v cmd/simple-server/main.go

data-creator:
	go build -o bin/data-creator -v cmd/data-creator/main.go