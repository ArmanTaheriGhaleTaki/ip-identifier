.PHONY: run build lint test
up: lint test 
	docker compose -f deployments/docker-compose.yml kill ; docker compose -f deployments/docker-compose.yml rm -f ; docker rmi ipidentifier:latest  ; docker compose -f deployments/docker-compose.yml  up  
run: 
	go run cmd/http_server/main.go
build:
	go build cmd/http_server/main.go
lint:
	golangci-lint run ./...
test:
	go test -v ./...
