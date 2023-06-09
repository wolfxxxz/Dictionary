build:
	go build -v ./cmd/slovarik/

run:
	./apiserver

migration_up:
	migrate -path migrations -database "postgres://localhost:5433/dictionary?sslmode=disable&user=postgres&password=1" up

migration_down:
	migrate -path migrations -database "postgres://localhost:5433/dictionary?sslmode=disable&user=postgres&password=1" down

commit:
	git commit -a -m "avtoSave"

push:
	git push origin master