.PHONY: run migrations migrate-up migrate-down tree

run:
	@air --build.cmd "go build -o ./.tmp/main cmd/main.go" --build.bin "./.tmp/main" --log.main_only "true"

migrations:
	@read -p "Enter the migration name: " name; \
	if [ -z "$$name" ] || echo "$$name" | grep -q '[^a-zA-Z0-9_]'; then \
		echo "Invalid migration name, it should only contain letters, numbers and underscores"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir migrations/sql -seq $$name

mgrup:
	@go run migrations/main.go up

mgrdown:
	@go run migrations/main.go down

cmpup:
	@docker compose up -d

cmpdown:
	@docker compose down

tree:
	@tree -a -I ".ropeproject|.git|tmp|migrations" | xclip -selection clipboard
