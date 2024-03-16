all: dev

CERTS_DIRECTORY = '.certs'

clean: down
	@echo "cleaning directory and volumes"
	yarn --cwd ./frontend/blog clean
	yarn --cwd ./backend/cms clean
	docker system prune -f --volumes

dev:
	@echo "Starting Dev"
	docker compose --profile development up -d

restart:
	docker compose --profile development restart

down:
	@echo "Shutting Down Dev"
	docker compose --profile development down -v

build:
	@echo "Building Dev Enviorment"
	docker compose --profile development up --build -d

stop:
	@echo "Stoping services"
	docker compose --profile development stop

test:
	@echo "Testing Qohelet"

deploy:
	@echo "Deploying Qohelet Backend"