MODULE_NAME=testcontainers-registry

VENDOR_DIR = vendor

GOLANGCI_LINT_VERSION ?= v1.49.0

GO ?= go
GOLANGCI_LINT ?= $(shell go env GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION)

goModules := $(shell find . -name 'go.mod' | xargs dirname | tail -n +2)
lintGoModules := $(subst -.,-module,$(subst /,-,$(addprefix lint-,$(goModules))))
tidyGoModules := $(subst -.,-module,$(subst /,-,$(addprefix tidy-,$(goModules))))
testGoModules := $(subst -.,-integration,$(subst /,-,$(addprefix test-,$(goModules))))

.PHONY: $(VENDOR_DIR)
$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)
	@$(GO) mod vendor
	@$(GO) mod tidy

.PHONY: lint
lint: $(GOLANGCI_LINT) $(VENDOR_DIR) $(lintGoModules)
	@$(GOLANGCI_LINT) run -c .golangci.yaml

.PHONY: $(lintGoModules)
$(lintGoModules): $(GOLANGCI_LINT)
	$(eval GO_MODULE := "$(subst lint/module,.,$(subst -,/,$(subst lint-module-,,$@)))")

	@echo ">> module: $(GO_MODULE)"
	@cd "$(GO_MODULE)"; $(GOLANGCI_LINT) run

.PHONY: tidy
tidy: $(tidyGoModules)
	@echo ">> tidy"
	@$(GO) mod tidy

.PHONY: $(tidyGoModules)
$(tidyGoModules):
	$(eval GO_MODULE := "$(subst tidy/module,.,$(subst -,/,$(subst tidy-module-,,$@)))")

	@echo ">> tidy module: $(GO_MODULE)"
	@cd "$(GO_MODULE)"; $(GO) mod tidy

.PHONY: test
test: test-integration $(testGoModules)

## Run integration tests
.PHONY: test-integration
test-integration:
	@echo ">> unit test"
	@$(GO) test -v -gcflags=-l -coverprofile=features.coverprofile -covermode=atomic -race ./...

.PHONY: $(testGoModules)
$(testGoModules):
	$(eval GO_MODULE := "$(subst test-integration-,,$@)")
	@echo ">> integration test: $(GO_MODULE)"
	@cd "$(GO_MODULE)"; $(GO) test -gcflags=-l -coverprofile=features.coverprofile -v ./...
	@echo

.PHONY: gha-vars
gha-vars:
	@echo "::set-output name=MODULE_NAME::$(MODULE_NAME)"
	@echo "::set-output name=GOLANGCI_LINT_VERSION::$(GOLANGCI_LINT_VERSION)"

$(GOLANGCI_LINT):
	@echo "$(OK_COLOR)==> Installing golangci-lint $(GOLANGCI_LINT_VERSION)$(NO_COLOR)"; \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin "$(GOLANGCI_LINT_VERSION)"
	@mv ./bin/golangci-lint $(GOLANGCI_LINT)
