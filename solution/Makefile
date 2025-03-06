migrate:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

test:
	go clean -testcache
	go test ./...

mock:
	go generate ./...

test.integration:
	go clean -testcache
	docker compose -f docker-compose.test.yml --env-file .env.test down
	docker compose -f docker-compose.test.yml --env-file .env.test up -d
	docker compose -f docker-compose.test.yml --env-file .env.test logs -f app-test > tests/logs/logs.txt 2>&1 &
	go test ./tests -v
	docker compose -f docker-compose.test.yml --env-file .env.test down

test:
	go clean -testcache
	go test --short -v ./...