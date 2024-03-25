all: dev

dev:
	@echo "Starting Dev"
	docker compose up -d

restart:
	docker compose restart

down:
	@echo "Shutting Down Dev"
	docker compose down -v

build:
	@echo "Building Dev Enviorment"
	docker compose up --build -d

stop:
	@echo "Stoping services"
	docker compose stop

test:
	@echo "Testing User"