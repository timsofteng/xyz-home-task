include .env

.DELETE_ON_ERROR:

APP_NAME=xyz-home-task-backend
HOST=0.0.0.0
HTTP_SERVER_PORT=8080
LOG_LEVEL=info
SPEC_PATH=../spec/openapi.yaml
HOST=0.0.0.0

set_env = export HOST=$(HOST) && \
          export APP_NAME=$(APP_NAME) && \
          export HTTP_SERVER_PORT=$(HTTP_SERVER_PORT) && \
          export LOG_LEVEL=$(LOG_LEVEL)

.PHONY:
build:
	CGO_ENABLED=0 go build -o bin/${APP_NAME} ./cmd/*

.PHONY:
app:
	$(set_env) && \
	./bin/${APP_NAME}

.PHONY: dev
dev:
	$(set_env) && \
	air

.PHONY:
run:
	$(set_env) && \
	go run ./cmd/*

.PHONY:
test:
	go test ./...

.PHONY:
lint:
	golangci-lint run --concurrency=2


.PHONY:
codegen-http-server-handler:
	oapi-codegen \
		-config ./spec/http-server.cfg.yaml \
		 ../spec/openapi.yaml

## CONTAINER #####################
.PHONY:
container-build:
	docker build \
		--pull \
		--build-arg APP_NAME=$(APP_NAME) \
		--build-arg HTTP_SERVER_PORT=$(HTTP_SERVER_PORT) \
		-t $(APP_NAME) .

.PHONY:
container-run:
	docker run \
		-e HOST=$(HOST) \
		-e HTTP_SERVER_PORT=$(HTTP_SERVER_PORT) \
		-it -p $(HTTP_SERVER_PORT):$(HTTP_SERVER_PORT) $(APP_NAME)
#########################################
