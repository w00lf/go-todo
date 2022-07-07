OUT ?= $(CURDIR)/dist
PROJECT ?=$(shell basename $(CURDIR))
SRC ?= $(CURDIR)
BINARY ?= $(OUT)/$(PROJECT)
SOURCES := $(shell find $(CURDIR) -type f -name '*.go' -not -name '*_test.go') go.mod go.sum
ifdef env_file
	include $(env_file)
	export
endif

all: run lint test bundle
.PHONY : all

build: $(BINARY)

$(BINARY): $(SOURCES) $(OUT)
	goreleaser build --single-target --output $(BINARY) --snapshot --rm-dist

run: $(BINARY)
	$(BINARY)

lint:
	golangci-lint run

test:
	go test

bundle:
	go mod tidy

