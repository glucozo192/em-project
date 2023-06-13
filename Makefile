postgres:
	docker run --name postgres2 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=6677028a -d postgres

createdb:
	docker exec -it postgres2 createdb --username=root --owner=root shop_vui

dropdb:
	docker exec -it postgres2 dropdb shop_vui

migrate-up:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/shop_vui?sslmode=disable" -verbose up

migrate-up1:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/shop_vui?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/shop_vui?sslmode=disable" -verbose down

migrate-down1:
	migrate -path db/migration -database "postgresql://root:6677028a@localhost:5431/shop_vui?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: createdb dropdb postgres migrate-up migrate-down sqlc server migrate-up1 migrate-down1 