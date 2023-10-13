postgres:
	docker run --name dqlgen-postgres -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=swejal -d postgres:12-alpine

migrateup:
	migrate -path db/migrations -database "postgresql://user:swejal@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://user:swejal@localhost:5432/postgres?sslmode=disable" -verbose down

start:
	docker start 469bc730a56f 

.PHONY: postgres migrateup migratedown start