SOURCE_DIR   := src
PROTO_DIR    := ../../../api/proto
SERVICE_DIR  := $(SOURCE_DIR)/internal/service
CLIENT_DIR   := $(SOURCE_DIR)/internal/client
MODULE_NAME  := github.com/FinnTew/FinnEcommerce
SERVICES     := auth cart checkout email order payment product shortlink user

.PHONY: gen_svr gen_client modtidy fmt lint clean test build

gen_svr: $(addprefix gen_svr-,$(SERVICES))

gen_svr-%:
	@echo "===> Generating service code for [$*]"
	cd $(SERVICE_DIR)/$* && \
	cwgo server \
		-I $(PROTO_DIR) \
		--type RPC \
		--module $(MODULE_NAME)/src/internal/service/$* \
		--service $* \
		--idl $(PROTO_DIR)/$*.proto

gen_client: $(addprefix gen_client-,$(SERVICES))

gen_client-%:
	@echo "===> Generating client code for [$*]"
	cd $(CLIENT_DIR) && \
	cwgo client \
		-I ../../api/proto \
		--type RPC \
		--module $(MODULE_NAME)/src/internal/client \
		--server_name $* \
		--idl ../../api/proto/$*.proto

modtidy: $(addprefix modtidy-,$(SERVICES))

modtidy-%:
	@echo "===> Running `go mod tidy` for [$*]"
	cd $(SERVICE_DIR)/$* && \
	go mod tidy

fmt:
	@echo "===> Formatting code..."
	cd $(SERVICE_DIR) && \
	for svc in $(SERVICES); do\
		echo "Formatting $$svc..."; \
		cd ./$$svc && go fmt ./...; \
		cd ../.; \
	done

lint:
	@echo "===> Running linters..."
	@echo "===> Running golangci-lint..."
	cd $(SERVICE_DIR) && \
	for svc in $(SERVICES); do \
		echo "Running golangci-lint for $$svc..."; \
		cd ./$$svc && golangci-lint run; \
		cd ../.; \
	done

clean:
	@echo "===> Cleaning generated files..."
	rm -rf $(SERVICE_DIR)/*/pb/*.go

test:
	@echo "===> Running tests..."
	cd $(SERVICE_DIR) && \
	for svc in $(SERVICES); do \
		echo "Running tests for $$svc..."; \
		cd ./$$svc && go test ./...; \
		cd ../.; \
	done

build:
	@echo "===> Building all services..."
	cd $(SOURCE_DIR)/cmd && \
	for svc in $(SERVICES); do \
		echo "Building $$svc..."; \
		cd./$$svc && sh build.sh; \
		cd../.; \
	done