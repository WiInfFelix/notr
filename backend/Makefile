
# Get the name of the root directory
ROOT_DIR := $(notdir $(CURDIR))
ATLAS_MIGRATE_DIR := $(CURDIR)/ent/migrate/migrations
POSTGRES_URL := postgres://postgres:postgres@localhost:5432/notr_db?sslmode=disable

build:
	go build -o $(ROOT_DIR)

run: build
	./$(ROOT_DIR)
	
clean:
	go mod tidy

test:
	go test -v

migrate:
	atlas migrate apply --dir "$(ATLAS_MIGRATE_DIR)" --url "$(POSTGRES_URL)"

migrate-create:
	atlas migrate diff --dir "$(ATLAS_MIGRATE_DIR)" --to "ent://ent/schema" --dev-url "docker://postgres/16/dev"

migrate-clean:
	rm -rf $(ATLAS_MIGRATE_DIR)/*