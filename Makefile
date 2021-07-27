.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

PATH = backend/migrations
PORT = 5432
DB = ecart
POSTGRESQL_URL = postgres://postgres:password@localhost:$(PORT)/$(DB)?sslmode=disable

migrate-create:
	/usr/bin/migrate create -ext sql -dir $(PATH) -seq ${MIGR}

migrate-up:
	/usr/bin/migrate -database $(POSTGRESQL_URL) -path $(PATH) up
migrate-down:
	/usr/bin/migrate -database $(POSTGRESQL_URL) -path $(PATH) down

