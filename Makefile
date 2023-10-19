all:
	PG_URL=postgres://user:pass@localhost:5433/translator \
	go run cmd/app/main.go
