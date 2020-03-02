.PHONY: up down

include .env
export

build-prod:
	docker-compose -f docker-compose.prod.yml build --no-cache

prod: build-prod
	docker-compose -f docker-compose.prod.yml up

prod-down:
	docker-compose -f docker-compose.prod.yml down

check-envs:
ifndef PG_PORT
	@echo Warning: PG_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif
ifndef REST_PORT
	@echo Warning: REST_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif

up: check-envs
	docker-compose up

down:
	docker-compose down

protoc:
	# example: make protoc path=domain/company/port/grpc/company.proto
	docker run --rm -v $(CURDIR):/app -w /app grpc/go protoc --go_out=plugins=grpc:. ${path}