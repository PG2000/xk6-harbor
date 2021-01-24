IMG ?= heww/k6

GOLANG=golang:1.15.6

SHELL := /bin/bash

GIT_HASH := $(shell git rev-parse --short=8 HEAD)

TMPDIR := $(shell mktemp -d /tmp/xk6-harbor.XXXX)

BUILDPATH=$(CURDIR)

BIN ?= $(CURDIR)/bin

$(BIN):
	mkdir -p "$(BIN)"

DOCKERCMD=$(shell which docker)

GO_MODIFY_TAGS_VERSION := 1.13.0
GO_MODIFY_TAGS := $(BIN)/gomodifytags-$(GO_MODIFY_TAGS_VERSION)

.PHONY: gomodifytags
gomodifytags:
	@$(GO_MODIFY_TAGS) > /dev/null 2>&1 \
	|| rm -f $(BIN)/gomodifytags-*
	@$(MAKE) $(GO_MODIFY_TAGS)

$(GO_MODIFY_TAGS):
	$(MAKE) $(BIN)
	cd "$(TMPDIR)" && \
	go mod init gomodifytags && \
	go get 'github.com/fatih/gomodifytags@v$(GO_MODIFY_TAGS_VERSION)' && \
	go build -mod=readonly -o $(GO_MODIFY_TAGS) github.com/fatih/gomodifytags && \
	cd -

SWAGGER_IMAGENAME := quay.io/goswagger/swagger
SWAGGER_VERSION := 0.27.0
SWAGGER=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH) ${SWAGGER_IMAGENAME}:v${SWAGGER_VERSION}

generate-client:
	- rm -rf pkg/harbor/{models,client}
	$(SWAGGER) generate client -f pkg/harbor/swagger.yaml --target pkg/harbor --template=stratoscale --additional-initialism=CVE --additional-initialism=GC

modify-tags: gomodifytags
	@for f in $(shell ls pkg/harbor/models/*.go); \
		do $(GO_MODIFY_TAGS) -file $${f} -all -add-tags js -transform camelcase --skip-unexported -w ; \
	done
	@for f in $(shell ls pkg/harbor/client/*/*_parameters.go); \
		do $(GO_MODIFY_TAGS) -file $${f} -all -add-tags js -transform camelcase --skip-unexported -w ; \
	done

go-generate: generate-client modify-tags

.PHONY: k6
k6:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -i -o k6 ./cmd/k6/main.go

build: go-generate k6

docker-build: build
	@echo "Beginning build $(GIT_HASH)"
	docker build --build-arg GOLANG=$(GOLANG) -f docker/Dockerfile -t $(IMG):$(GIT_HASH) .
	docker tag $(IMG):$(GIT_HASH) $(IMG):latest

test: go-generate
	go test ./...

# Run go vet against code
.PHONY: vet
vet: go-generate
	go vet ./...

# find or download golangci-lint
# download golangci-lint if necessary
GOLANGCI_LINT := $(BIN)/golangci-lint
GOLANGCI_LINT_VERSION := 1.36.0

.PHONY: golangci-lint
golangci-lint:
	@$(GOLANGCI_LINT) version --format short 2>&1 \
		| grep '$(GOLANGCI_LINT_VERSION)' > /dev/null \
	|| rm -f $(GOLANGCI_LINT)
	@$(MAKE) $(GOLANGCI_LINT)

$(GOLANGCI_LINT):
	$(MAKE) $(BIN)
	# https://golangci-lint.run/usage/install/#linux-and-windows
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(BIN) 'v$(GOLANGCI_LINT_VERSION)'

# Run go linters
.PHONY: go-lint
go-lint: golangci-lint vet go-generate
	$(GOLANGCI_LINT) run --verbose --max-same-issues 0 --sort-results -D wrapcheck -D exhaustivestruct -D errorlint -D goerr113 -D gomnd -D nestif -D funlen -D gosec

clean:
	- rm -rf pkg/harbor/{models,client}
	- rm -rf $(BIN)
