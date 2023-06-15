BUILD_VERSION=dev-snapshot

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(dir $(MAKEFILE_PATH))

DOCKER_COMPOSE_FILE_PATH := -f test/system/docker-compose.yaml -f test/system/docker-compose.local.yaml
DOCKER_COMPOSE_ENV_FILE_PATH := test/system/.env

# https://docs.github.com/en/actions/learn-github-actions/variables
ifeq (${CI},true)
    DOCKER_COMPOSE_FILE_PATH := -f test/system/docker-compose.yaml
endif

# Loading docker .env file if present
ifneq (,$(wildcard $(DOCKER_COMPOSE_ENV_FILE_PATH)))
	include $(DOCKER_COMPOSE_ENV_FILE_PATH)
    export
endif

# No API key is being set
ifeq (${DD_API_KEY},)
    DOCKER_COMPOSE_FILE_PATH := -f test/system/docker-compose.yaml
endif

all: build

.PHONEY: build
build: ## Build the application
	cd cmd && go build -ldflags="-X pkg/config.BuildVersion=$(BUILD_VERSION)" -o ../bin/kubehound kubehound/*.go

.PHONY: infra-rm
infra-rm: ## Delete the testing stack
	docker compose $(DOCKER_COMPOSE_FILE_PATH) rm -fvs 

.PHONY: infra-up
infra-up: ## Spwan the testing stack
	docker compose $(DOCKER_COMPOSE_FILE_PATH) up -d

.PHONY: test
test: ## Run the full suite of unit tests 
	$(MAKE) infra-rm
	$(MAKE) infra-up
	cd pkg && go test ./...

.PHONY: system-test
system-test: ## Run the system tests
	# $(MAKE) infra-rm
	# $(MAKE) infra-up
	cd test/system && export KUBECONFIG=/home/edouard/dd/KubeHound/test/setup/.kube/config && go test -v -timeout "60s" -race ./...

.PHONY: local-cluster-reset
local-cluster-reset: ## Destroy the current kind cluster and creates a new one
	$(MAKE) local-cluster-destroy
	$(MAKE) local-cluster-create
	$(MAKE) local-cluster-config-deploy

.PHONY: local-cluster-deploy
local-cluster-deploy: ## Create a kind cluster with some vulnerables resources (pods, roles, ...)
	$(MAKE) local-cluster-create
	$(MAKE) local-cluster-config-deploy

.PHONY: local-cluster-config-deploy
local-cluster-config-deploy: ## Deploy the attacks resources
	bash test/setup/manage-cluster-resources.sh deploy

.PHONY: local-cluster-config-delete
local-cluster-config-delete: ## Delete the attack resources
	bash test/setup/manage-cluster-resources.sh delete

.PHONY: local-cluster-create
local-cluster-create: ## Create a local kind cluster without any data
	bash test/setup/manage-cluster.sh create

.PHONY: local-cluster-destroy
local-cluster-destroy: ## Destroy the local kind cluster
	bash test/setup/manage-cluster.sh destroy

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'