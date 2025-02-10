postgres:
	docker run --name short-url-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it short-url-db createdb --username=root --owner=root shorturldb

dropdb:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/shorturldb?sslmode=disable" -verbose drop

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/shorturldb?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/shorturldb?sslmode=disable" -verbose down

run:
	go run ./cmd/main.go

.PHONY: postgres createdb dropdb migrateup migratedown