APP_NAME=xyz-home-task-api-doc
SERVER_PORT=3030

.PHONY:
dev:
	bun dev

.PHONY:
lint:
	bun lint:openapi


## CONTAINER #####################
.PHONY:
container-build:
	docker build \
    --pull \
		-t $(APP_NAME) .

.PHONY:
container-run:
	docker run \
		-it -p ${SERVER_PORT}:80 $(APP_NAME)
#########################################
