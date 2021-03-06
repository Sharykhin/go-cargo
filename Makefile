.PHONY: up down

include .env
export

check-envs:
ifndef PG_PORT
	@echo Warning: PG_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif
ifndef REST_PORT
	@echo Warning: REST_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif
ifndef GRPC_COMPANY_PORT
	@echo Warning: GRPC_COMPANY_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif

up-grpc: check-envs
	docker-compose up postgres grpc-company

up: check-envs
	docker-compose up

down:
	docker-compose down

protoc:
	# example: make protoc path=domain/company/port/grpc/company.proto
	docker run --rm -v $(CURDIR):/app -w /app grpc/go protoc --go_out=plugins=grpc:. ${path}

migration:
	# example: make migration name=crate_company_table
	docker-compose run sql-migration goose -dir infrastructure/database/migration create ${name} sql

migrate-up:
	# example: make migrate-up
	docker-compose run sql-migration goose -dir infrastructure/database/migration postgres "host=postgres user=postgres password=root dbname=gocargo sslmode=disable port=5432" up

migrate-down:
	# example: make migrate-down
	docker-compose run sql-migration goose -dir infrastructure/database/migration postgres "host=postgres user=postgres password=root dbname=gocargo sslmode=disable port=5432" down

migrate-status:
	# example: make migrate-status
	docker-compose run sql-migration goose -dir infrastructure/database/migration postgres "host=postgres user=postgres password=root dbname=gocargo sslmode=disable port=5432" status

grpc-cli-ls:
	# read more: https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	# example: make grpc-cli-ls path="company.company" port=50052
	docker run -v $(CURDIR):/defs --rm -it --net go_cargo namely/grpc-cli ls docker.for.mac.localhost:${port} ${path}

grpc-cli-call:
	# example: make grpc-cli-call path='CreateCompany ""' port=50052
	docker run -v $(CURDIR):/defs --rm -it --net go_cargo namely/grpc-cli call docker.for.mac.localhost:${port} ${path}