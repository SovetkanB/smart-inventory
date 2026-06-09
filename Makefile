DB_URL=postgres://postgres:postgres@localhost:5432/smart_inventory?sslmode=disable

MIGRATIONS_DIR=./migrations

migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-down-all:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down -all

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version