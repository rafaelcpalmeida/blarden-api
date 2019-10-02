APP_NAME := blarden-api
API_PORT := 8080
DOCKER_BUILD_IMAGE := rafaelcpalmeida/blarden-api

docker-build:
	@docker build -f Dev.Dockerfile . -t $(DOCKER_BUILD_IMAGE)

docker-run:
	@docker run -it --rm --network "own-network" \
	-v "$$(pwd):/app" \
	-v ".go" \
	-p $(API_PORT):$(API_PORT) \
	-e PORT=$(API_PORT)	\
	-e DEBUG="true" \
	-e POSTGRESQL_DBNAME="blarden-api_development" \
	-e POSTGRESQL_HOST="postgresql" \
	-e POSTGRESQL_USER="postgres" \
	-e POSTGRESQL_PASSWORD="postgres" \
	-e POSTGRESQL_PORT="5432" \
	--name $(APP_NAME) \
	-w /app $(DOCKER_BUILD_IMAGE)

build:
	@echo "Making $(APP_NAME)"
	make docker-build

run:
	@echo "Running application API $(APP_NAME)"
	make docker-run

test:
	@echo "Testing application $(APP_NAME)"
	@docker run -it --rm \
	-v "$$(pwd):/app" \
	-v ".go" \
	-e DEBUG="true" \
	-e TESTING="true" \
	--name $(APP_NAME) \
	-w /app $(DOCKER_BUILD_IMAGE) sh -c "go test ./..."
