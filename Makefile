tidy:
	@go mod tidy
	@go mod vendor
run:
	@go run cmd/main.go
seed:
	@go run cmd/main.go --seed
migrate:
	@go run cmd/main.go --migrate
migrate-seed:
	@go run cmd/main.go --migrate --seed
migrate-fresh:
	@go run cmd/main.go --migrate --fresh