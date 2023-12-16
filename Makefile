include .env

createdb:
	docker run --name database -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -p 5432:5432 -d postgres:12-alpine

dropdb:
	docker exec -it database dropdb $(DB_NAME)