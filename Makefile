ROOT := .
TMP_DIR := tmp
BIN := $(TMP_DIR)/main
DOCKERFILE_DIR := .
ENVFILE_DIR := .env.local

.PHONY: build run clean github docker-br docker-rm

# Require Ubuntu
build:
	@echo "Building..."
	@mkdir -p $(TMP_DIR)
	go build -o $(BIN) ./cmd/instay

run: build
	@echo "Running..."
	@$(BIN)

clean:
	@echo "Cleaning..."
	@rm -rf $(TMP_DIR)

# Require Windows
github:
	@if "$(CM)"=="" ( \
		echo Usage: make github CM="commit message" && exit 1 \
	)
	git add .
	git commit -m "$(CM)"
	git push
	git push clone

# Require Docker
docker-br:
	docker build -t instay-be $(DOCKERFILE_DIR)
	docker run --env-file $(ENVFILE_DIR) -d -p 8080:8080 --name instay instay-be

docker-rm:
	docker stop instay 
	docker rm instay
	docker rmi instay-be