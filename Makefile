MAIN_FILE=cmd/main.go

run:
	go run $(MAIN_FILE) start

# Usage: make create-migration name=<migration_name>
create-migration:
	go run $(MAIN_FILE) create-migration $(name)


# Usage: run the migration 
migrate:
	go run $(MAIN_FILE) migrate

# Usage: make rollback version=<migration_version>
rollback:
	go run $(MAIN_FILE) rollback $(version)