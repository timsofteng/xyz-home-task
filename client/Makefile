APP_NAME=xyz-home-task-frontend

HOST=0.0.0.0
SERVER_PORT=3000
API_PORT=8080

# .DELETE_ON_ERROR:

.PHONY:
check-deps:
	inkit-check-deps requirements

.PHONY:
build:
	export \
		VITE_HOST=$(HOST) \
		VITE_SERVER_PORT=$(SERVER_PORT) \
		VITE_API_PORT=$(API_PORT) && \
		pnpm build

.PHONY: dev
dev:
	export \
		VITE_HOST=$(HOST) \
		VITE_SERVER_PORT=$(SERVER_PORT) \
		VITE_API_PORT=$(API_PORT);  \
		pnpm dev

.PHONY:
test:
	export \
		VITE_HOST=$(HOST) \
		VITE_SERVER_PORT=$(SERVER_PORT) \
		VITE_API_PORT=$(API_PORT) && \
		pnpm test

.PHONY:
lint:
	pnpm lint

.PHONY:
codegen-openapi-client:
	pnpm openapi-client-gen

## CONTAINER #####################
.PHONY:
container-build:
	docker build \
    --pull \
		--build-arg HOST=$(HOST) \
		--build-arg SERVER_PORT=$(SERVER_PORT) \
		--build-arg API_PORT=$(API_PORT) \
		-t $(APP_NAME) .

.PHONY:
container-run:
	docker run \
		-it -p ${SERVER_PORT}:80 $(APP_NAME)
#########################################
