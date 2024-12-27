up:lint
	docker compose -f deployments/docker-compose.yml kill ; docker compose -f deployments/docker-compose.yml rm -f ; docker rmi ipidentifier:latest  ; docker compose -f deployments/docker-compose.yml  up  
run: 
	go run cmd/main.go
build:
	go build cmd/main.go
lint:
	golangci-lint run ./...

