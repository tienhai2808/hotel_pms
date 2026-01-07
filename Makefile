ROOT := .
TMP_DIR := tmp
BIN := $(TMP_DIR)/main

.PHONY: build run clean github

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