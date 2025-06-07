.PHONY: up down build dev migrate migrate-new migrate-apply migrate-down clean generate

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

migrate:
	atlas migrate status --url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo"

migrate-new:
	atlas migrate new --url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo" --dev-url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo"

migrate-apply:
	atlas migrate apply --url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo"

migrate-down:
	atlas migrate down --url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo" --dev-url "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable&search_path=todo"

clean:
	docker compose down -v
	rm -rf tmp/* 

tbls-generate:
	docker compose --profile tbls run tbls

generate:
	oapi-codegen -config api/config.yaml api/openapi.yaml
