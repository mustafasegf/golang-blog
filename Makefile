pullpostgres:
	docker pull postgres:12-alpine

postgres:
	docker run --name golang-blog -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

removecontainer:
	docker rm golang-blog

createdb:
	docker exec -it golang-blog createdb --username=root --owner=root golang-blog

dropdb:
	docker exec -it golang-blog dropdb golang-blog

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/golang-blog?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/golang-blog?sslmode=disable" -verbose down

# note: used only in dev
sqlc:
	sqlc generate

build-nix:
	go build -o ./bin main.go; ./bin/main

build-win:
	go build -o .\bin main.go; .\bin\main

watch:
	air -c .air.conf
.PHONY: pullpostgres postgres removecontainer createdb dropdb migrateup migratedown sqlc build-nix buiid-win watch