SHELL := /bin/bash

TARGET := hetzner-sb-notifier
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

.PHONY: all build clean uninstall fmt simplify check run

all: check build

$(TARGET):
	@go build -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@echo "Performing clean"
	@rm -f $(TARGET)

uninstall: clean
	@echo "Performing uninstall"
	@rm -f $$(which ${TARGET})

fmt:
	@echo "Performing fmt"
	@gofmt -l -w .

simplify:
	@echo "Performing simplify"
	@gofmt -s -l -w .

check:
	@echo "Performing check"
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./...); do golint $${d}; done
	@go tool vet .

run: install
	@echo "Performing run"
	@$(TARGET)