.PHONY: up down

up:
	docker-compose up

down:
	docker-compose down

build-prod:
	docker-compose -f docker-compose.prod.yml build --no-cache

prod: build-prod
	docker-compose -f docker-compose.prod.yml up

prod-down:
	docker-compose -f docker-compose.prod.yml down