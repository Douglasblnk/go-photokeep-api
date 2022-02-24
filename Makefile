generate:
	go run github.com/99designs/gqlgen generate

dev:
	go run github.com/cosmtrek/air -c ./cmd/.air.toml
