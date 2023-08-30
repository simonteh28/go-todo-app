#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up: ## Run migrations UP
migrate-up:
	docker-compose --profile tools run --rm migrate up

migrate-down: ## Rollback migrations
migrate-down:
	docker-compose --profile tools run --rm migrate down 1

migrate-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-create:
	docker-compose --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

console-db: ## Enter to database console
console-db:
	docker-compose exec -it db psql -U postgres -d postgres