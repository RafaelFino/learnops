default:
	go build -o bin/simple-server -v cmd/simple-server/main.go
	go build -o bin/data-creator -v cmd/data-creator/main.go
	go build -o bin/get-currency -v cmd/get-currency/main.go
	go build -o bin/db-connect -v cmd/db-connect/main.go	
	go build -o bin/db-exec -v cmd/db-exec/main.go	
	go build -o bin/db-api -v cmd/db-api/main.go	
	go build -o bin/chat-server -v cmd/chat-server/main.go	
	go build -o bin/chat-client -v cmd/chat-client/main.go	

simple-server:
	go build -o bin/simple-server -v cmd/simple-server/main.go

data-creator:
	go build -o bin/data-creator -v cmd/data-creator/main.go

get-currency:
	go build -o bin/get-currency -v cmd/get-currency/main.go	

db-connect:
	go build -o bin/db-connect -v cmd/db-connect/main.go		

db-execute:
	go build -o bin/db-exec -v cmd/db-exec/main.go	

db-api:
	go build -o bin/db-api -v cmd/db-api/main.go	

chat:
	go build -o bin/chat-server -v cmd/chat-server/main.go	
	go build -o bin/chat-client -v cmd/chat-client/main.go	