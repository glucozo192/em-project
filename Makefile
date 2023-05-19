postgres:
	docker run --name postgres2 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=6677028a -d postgres

createdb:
	docker exec -it postgres2 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres2 dropdb simple_bank

migrate-up:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/simple_bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
.PHONY: createdb dropdb postgres migrate-up migrate-down sqlc