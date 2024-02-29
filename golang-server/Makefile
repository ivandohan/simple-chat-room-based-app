postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgresinit-L:
	sudo docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root chat-room-app

dropdb:
	docker exec -it postgres15 dropdb chat-room-app

migrate-init:
	migrate create -ext sql -dir database/migrations/ -seq init_migration

migrate-up:
	sudo migrate -path database/migrations/ -database postgresql://root:password@localhost:5433/chat-room-app?sslmode=disable -verbose up

migrate-rollback:
	sudo migrate -path database/migrations/ -database postgresql://root:password@localhost:5433/chat-room-app?sslmode=disable -verbose down

.PHONY: postgresinit postgresinit-L postgres createdb dropdb migrate-init migrate-up migrate-rollback