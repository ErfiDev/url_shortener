init:
	go mod download

run:
	go run main.go

run_dependencies:
	docker compose up -d