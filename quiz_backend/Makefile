postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it postgres createdb --username=root --owner=root db

dropdb:
	docker exec -it postgres dropdb db

migration:
	migrate create -ext sql -dir db/migration -seq schema1

migrateUp:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/db?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/db?sslmode=disable" -verbose down

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl 

server:
	go run main.go

.PHONY: proto createdb dropdb migration migrateUp migrateDown evans server