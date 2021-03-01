.DEFAULT_GOAL := help

.PHONY: gen
gen: proto/gen ## gernerate codes
	go generate ./...

.PHONY: lint
lint: proto/lint ## lint codes
	golangci-lint run

.PHONY: fmt
fmt: proto/fmt ## format codes
	goimports -w .

.PHONY: proto/gen
proto/gen:
	@for f in example/**/**/**/**/*.proto; do \
		protoc \
			--proto_path=.:. \
			--proto_path=.:${GOPATH}/src \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--gotemplate_out=destination_dir=.,template_dir=./templates:. \
			$$f; \
		echo $$f; \
	done

.PHONY: proto/lint
proto/lint:
	buf check lint

.PHONY: proto/fmt
proto/fmt:
	prototool format -d . || true
	prototool format -w .

.PHONY: __
__:
	@echo "\033[33m"
	@echo "github.com/kzmake/dapr-kit"
	@echo "\033[0m"

.PHONY: help
help: __
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'
