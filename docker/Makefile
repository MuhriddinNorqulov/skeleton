docker-run.build:
	@echo "Running Docker container..."
	@docker compose -f docker-compose.build.yml --env-file ./../.env up --build -d


docker-run.dev:
	@echo "Running Docker container..."
	@docker compose -f docker-compose.dev.yml --env-file ./../.env up -d

docker-run.local:
	@echo "Running Docker container..."
	@docker compose -f docker-compose.local.yml --env-file ./../.env up --build -d
