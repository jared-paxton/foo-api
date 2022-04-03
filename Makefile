build:
	go build -o bin/foo-api-server cmd/server/main.go

run:
	go run cmd/server/main.go

clean:
	rm bin/foo-api-server

