pwd := $(shell pwd)

all: cpp go doc

.PHONY: cpp
cpp:
	@echo "Generating apm-library package for C++"
	@docker run --user `id -u` --rm -v $(PWD):/defs namely/protoc-all:1.51_1 -d . -l cpp -o cpp

.PHONY: go
go:
	@echo "Generating apm-library package for Go"
	@docker run --user `id -u` --rm -v $(PWD):/defs namely/protoc-all:1.51_1 -d . -l go -o go
	@cd go/collectorpb
	@go mod init github.com/solarwinds/apm-proto/tree/main/go/collectorpb
	@go mod tidy

.PHONY: doc
doc:
	@echo "Generating README.md"
	@docker run --rm -v "${PWD}":/out -v "${PWD}":/protos pseudomuto/protoc-gen-doc --doc_opt=markdown,README.md	

.PHONY: check
check:	all
	git diff --exit-code

clean:
	@echo "Cleaning packages"
	@rm -rf ./cpp
	@rm -rf ./go
	@echo "Cleaning markdown files"
	@rm -rf *.md
