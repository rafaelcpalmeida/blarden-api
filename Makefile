APP_NAME := blarden-api
API_PORT := 8080
DOCKER_BUILD_IMAGE := rafaelcpalmeida/blarden-api

# Gold <3
# https://news.ycombinator.com/item?id=11939200
.PHONY: help
help:	### this screen. Keep it first target to be default
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

.PHONY: docker-build
docker-build: ### build development docker image
	@echo "Buildingin $(APP_NAME)"
	@docker build -f Dev.Dockerfile . -t $(DOCKER_BUILD_IMAGE)

.PHONY: run
run: ### run previously builded image
	@echo "Running $(APP_NAME)"
	@docker run -it --rm --network "blarden-network" \
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

test: ### run application test suite
	@echo "Testing $(APP_NAME)"
	@docker run -it --rm \
	-v "$$(pwd):/app" \
	-v ".go" \
	-e DEBUG="true" \
	-e TESTING="true" \
	--name $(APP_NAME) \
	-w /app $(DOCKER_BUILD_IMAGE) sh -c "go test -cover ./..."
