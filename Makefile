include .env

to-dev:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	docker-compose up --scale worker=2 -d --build
	$(MAKE) migrate_up
	$(MAKE) run_seeder

run:
	docker-compose up --scale worker=2 -d --build
	go run cmd/app/main.go

create_migration:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose create $(name) sql

migrate_up:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose up

migrate_down:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose down

run_seeder:
	go run cmd/seed/seeder.go

resharding:
	# Установка wal_level на master-узле
	docker compose exec -T master psql -U postgres -c "ALTER SYSTEM SET wal_level = 'logical';"
	docker compose exec -T master psql -U postgres -c "SELECT run_command_on_workers('ALTER SYSTEM SET wal_level = ''logical''');"

	# Перезапуск всех контейнеров
	docker compose stop && docker compose start
	sleep 5

	# Запуск rebalance
	docker compose exec -T master psql -U postgres -c "SELECT citus_rebalance_start();"