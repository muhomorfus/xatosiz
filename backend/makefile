MIGRATIONS_POSTGRES_DIR := deployments/migrations/postgres
MIGRATIONS_CLICKHOUSE_DIR := deployments/migrations/clickhouse

POSTGRES := postgres://user:password@127.0.0.1:5435/backend?sslmode=disable
CLICKHOUSE := tcp://localhost:9000/xatosiz?username=xatosiz&password=password&dial_timeout=500ms
KAFKA_BROKERS := broker:9092

POSTGRES := postgres://xatosiz:270fAmE2MGUiNsWoimMZmDOGGOUUg4U@90.156.209.107:5432/xatosiz?sslmode=disable
CLICKHOUSE := tcp://185.154.195.95:9000/xatosiz?username=xatosiz&password=fVcB9enLiAK6hee8wlL4GieGWCGcbmhD&dial_timeout=500ms
KAFKA_BROKERS := 172.16.16.10:9092

.PHONY: migrate-postgres
migrate-postgres:
	@goose --dir=$(MIGRATIONS_POSTGRES_DIR) postgres "$(POSTGRES)" up

.PHONY: migrate-down-postgres
migrate-down-postgres:
	@goose --dir=$(MIGRATIONS_POSTGRES_DIR) postgres "$(POSTGRES)" down

.PHONY: build-goose-custom
build-goose-custom:
	go build -o bin/goose cmd/migrate/main.go

.PHONY: migrate-clickhouse
migrate-clickhouse: build-goose-custom
	./bin/goose -dir $(MIGRATIONS_CLICKHOUSE_DIR) -brokers $(KAFKA_BROKERS) -dsn "$(CLICKHOUSE)" -cmd up

.PHONY: migrate-down-clickhouse
migrate-down-clickhouse: build-goose-custom
	./bin/goose -dir $(MIGRATIONS_CLICKHOUSE_DIR) -brokers $(KAFKA_BROKERS) -dsn "$(CLICKHOUSE)" -cmd down

MIGRATION := default

.PHONY: new-migration-postgres
new-migration-postgres:
	goose --dir="$(MIGRATIONS_POSTGRES_DIR)" create $(MIGRATION) sql

.PHONY: new-migration-clickhouse
new-migration-clickhouse:
	goose --dir="$(MIGRATIONS_CLICKHOUSE_DIR)" create $(MIGRATION) go

.PHONY: build-technical
build-technical:
	mkdir -p bin
	go build -o bin/technical cmd/technical/main.go

.PHONY: build-web
build-web:
	mkdir -p bin
	go build -o bin/web cmd/web/main.go

.PHONY: .generate-swagger
.generate-swagger:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api/swagger.yml \
      -g go-server \
      -o /local/internal/openapi/ \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

.PHONY: goimports
goimports:
	goimports -w .

.PHONY: generate-swagger
generate-swagger: .generate-swagger goimports

TAG := lownie/services:xatosiz-backend-$(shell date +%Y%m%d%H%M%S)

docker-push:
	@docker build --tag $(TAG)  .
	@docker push $(TAG)
	@echo Service pushed with tag: $(TAG)
