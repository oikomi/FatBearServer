GO      = go
BIN      = bin


MODULE   = $(shell $(GO) list -m)


all: fmt 
	$(GO) build \
		-tags release \
		-o $(BIN)/$(basename $(MODULE))  cmd/main.go


fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...