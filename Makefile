build: 
	go build -o bin/chainStore

run: build
	./bin/chainStore

postgres:
	 docker run --name furnitureShop -p 5433:5432 -e POSTGRES_USER=nurken -e POSTGRES_PASSWORD=123nura123 -d postgres:18.0-alpine3.22

createdb:
	docker exec -it furnitureShop createdb --username=nurken --owner=nurken furnitureShop 

dropdb:
	docker exec -it furnitureShop dropdb --username=nurken furnitureShop 

migrateup:
	migrate -path db/migration -database "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable" -verbose down

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nurkenti/furnitureShop/db/sqlc Store

coverage:
	go test -coverprofile=coverage.out ./api/...
	go tool cover -func=coverage.out


.PHONY: run, postgres, createdb, dropdb, migrateup, migratedowm, test, build, server, mock, coverage

test: 
	go test -v ./...

