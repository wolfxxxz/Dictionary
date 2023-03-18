build:
	go build -v ./cmd/apiserver/

run:
	./apiserver

migration_up:
	migrate -path migrations -database "postgres://localhost:5433/dictionary?sslmode=disable&user=postgres&password=1" up

migration_down:
	migrate -path migrations -database "postgres://localhost:5433/dictionary?sslmode=disable&user=postgres&password=1" down