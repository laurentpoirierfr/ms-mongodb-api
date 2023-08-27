
test:
	go test -v -cover -short ./...

mongodb:
	docker-compose up

server: mongodb
	go run ./cmd/server/main.go

swagger:
	swag init -g cmd/server/main.go
	
.PHONY: mongodb server test swagger
