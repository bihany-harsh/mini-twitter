postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:16-alpine
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root mini-twitter
dropdb:
	docker exec -it postgres16 dropdb mini-twitter
migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/mini-twitter?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/mini-twitter?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc