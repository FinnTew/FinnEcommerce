SOURCE_DIR   := src
PROTO_DIR    := ../../../api/proto
SERVICE_DIR  := $(SOURCE_DIR)/internal/service

MODULE_NAME  := github.com/FinnTew/FinnEcommerce

GENERATE_SERVICES     := auth cart checkout email order payment product shortlink user
TIDY_SERVICES     := auth cart checkout email order payment product shortlink user

.PHONY: all
all: generate

.PHONY: generate
generate: $(GENERATE_SERVICES)

$(GENERATE_SERVICES):
	@echo "===> Generating service code for [$@]"
	cd $(SERVICE_DIR)/$@ && \
	cwgo server \
		-I $(PROTO_DIR) \
		--type RPC \
		--module $(MODULE_NAME)/src/internal/service/$@ \
		--service $@ \
		--idl $(PROTO_DIR)/$@.proto

.PHONY: modtidy
modtidy: $(TIDY_SERVICES)

$(TIDY_SERVICES):
	@echo "===> Run `go mod tidy` for [$@]"
	cd $(SERVICE_DIR)/$@ && \
	go mod tidy


.PHONY: fmt
fmt:
	@echo "===> Formatting code..."
	go fmt ./...

.PHONY: lint
lint:
	@echo "===> Running linters..."
	@echo "===> Running golangci-lint..."
	golangci-lint run ./...

.PHONY: clean
clean:
	@echo "===> Cleaning generated files..."
	 rm -rf $(SERVICE_DIR)/*/pb/*.go

.PHONY: test
test:
	@echo "===> Running tests..."
	go test ./... -v

.PHONY: build
build:
	@echo "===> Building all services..."
	for svc in $(SERVICES); do \
		echo "Building $$svc..."; \
		cd $(SOURCE_DIR)/cmd/$$svc && go build -o $$svc main.go; \
	done