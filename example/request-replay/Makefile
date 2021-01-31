.DEFAULT_GOAL := help

.PHONY: generate
generate: proto/generate ## gernerate codes
	go generate ./...

.PHONY: lint
lint: proto/lint ## lint codes
	golangci-lint run

.PHONY: fmt
fmt: proto/fmt ## format codes
	goimports -w .

.PHONY: proto
proto/generate:
	protoc \
	--proto_path=.:. \
	--proto_path=.:${GOPATH}/src \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--gotemplate_out=destination_dir=.,template_dir=../../template:. \
	./api/alice/v1/*.proto

	protoc \
	--proto_path=.:. \
	--proto_path=.:${GOPATH}/src \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--gotemplate_out=destination_dir=.,template_dir=../../template:. \
	./api/bob/v1/*.proto

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
	@echo "request-replay"
	@echo "\033[0m"

.PHONY: help
help: __
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'
