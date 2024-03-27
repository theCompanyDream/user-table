all: dev

ifeq ($(wildcard .env),)
    # .env file does not exist
	@echo ".env file does not exist. Creating one."
	cp .env.example .env
	@echo "Created .env file."
else
	# .env file exists
	@echo ".env file exists."
endif

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