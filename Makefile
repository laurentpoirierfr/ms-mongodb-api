
test:
	go test -v -cover -short ./...

mongodb:
	docker-compose up

swagger:
	swag init -g cmd/server/main.go -o ./api

server: mongodb swagger
	go run ./cmd/server/main.go

.PHONY: mongodb server test swagger
