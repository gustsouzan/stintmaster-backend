docker/up:
	docker compose --env-file .docker-compose.env -f docker-compose.yml up -d
 
docker/down:
	docker compose --env-file .docker-compose.env down -v

db/postgres/migrate/test-new:
	goose postgres -dir ./integrations/postgres/migrations postgres create $(name) sql	