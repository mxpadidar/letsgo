.PHONY: pgstart pgstop pgconnect migrations migrate-up migrate-down setup run

POSTGRES_CONTAINER_NAME = letsgo-db
POSTGRES_PASSWORD = secret
POSTGRES_USER = letsgo
POSTGRES_DB = letsgo
POSTGRES_PORT = 5432
POSTGRES_VOLUME = .tmp/pgdata
POSTGRES_IMAGE = docker.io/library/postgres:17-bookworm
POSTGRES_URL = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

pgstop:
	@podman stop "$(POSTGRES_CONTAINER_NAME)"
	@podman rm "$(POSTGRES_CONTAINER_NAME)"

pgstart: pgstop
	@mkdir -p "$(POSTGRES_VOLUME)"
	@podman run -d \
		--name "$(POSTGRES_CONTAINER_NAME)" \
		-e POSTGRES_PASSWORD="$(POSTGRES_PASSWORD)" \
		-e POSTGRES_USER="$(POSTGRES_USER)" \
		-e POSTGRES_DB="$(POSTGRES_DB)" \
		-p $(POSTGRES_PORT):5432 \
		-v "./$(POSTGRES_VOLUME):/var/lib/postgresql/data" \
		"$(POSTGRES_IMAGE)"

pgconnect:
	podman exec -it "$(POSTGRES_CONTAINER_NAME)" psql -U "$(POSTGRES_USER)" -d "$(POSTGRES_DB)"

migrations:
	@read -p "enter migration filename (do not use spaces and special characters): " filename; \
	goose -dir migrations create $$filename sql

migrate-up:
	@goose -dir migrations postgres $(POSTGRES_URL) up

migrate-down:
	@goose -dir migrations postgres $(POSTGRES_URL) down

run:
	@air --build.cmd "go build -o ./.tmp/main cmd/main.go" --build.bin "./.tmp/main" --log.silent="true" --tmp_dir=".tmp"
