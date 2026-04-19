dev: 
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build

up:
	docker compose up -d --build
down:
	docker compose down