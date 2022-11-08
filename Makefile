.DEFAULT_GOAL := help
DC := docker-compose -f deployments/docker-compose.yaml

.PHONY: help
help:
	@grep -E '(^[0-9a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

##
##Infra
##-------

.PHONY: start
start: ## Start docker containers.
	${DC} start redis db

.PHONY: stop
stop: ## Stop docker containers.
	${DC} stop

.PHONY: ps
ps: ## Show running containers.
	${DC} ps

.PHONY: rebuild
recreate: ## Recreate containers.
	${DC} down
	${DC} up -d redis db
	make ps

.PHONY: helm_dry_run
helm_dry_run: ## HELM dry run.
	helm upgrade --dry-run --debug -n backend --install --set image.repository=test --set image.tag=0.0.1 -f ./deployments/helm/values.yaml backend ./deployments/helm

##
##Ap  doplication
##-------

.PHONY: bot
bot:
	go run cmd/*.go bot run

.PHONY: server
server:
	go run cmd/*.go server run

.PHONY: bot_send
bot_send:
	go run cmd/*.go bot send

.PHONY: migrations_migrate
migrations_migrate: ## Run migrations.
	go run cmd/*.go migrations migrate

.PHONY: migrations_create
migrations_create:
	go run cmd/*.go migrations create ${n}

.PHONY: update_api
update_api_client: ## Update backend-api-client package.
	go get -u gitlab.com/ask-crew/clients/go/backend-api-client
	go mod tidy
