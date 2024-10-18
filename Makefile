start:
	docker-compose exec api go run cmd/api/main.go
	exit 0

up:
	docker-compose up -d

down:
	docker-compose down --timeout 0

deps:
	go mod tidy
