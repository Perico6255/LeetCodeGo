MODULE := github.com/Perico6255/LeetCodeGo

PKGS ?= ./...

TESTARGS ?=

COVERAGE_FILE := coverage.out

.PHONY: all fmt vet lint test test-verbose test-single coverage clean

all: fmt vet lint test

fmt:
	@echo "→ go fmt"
	go fmt $(PKGS)

vet:
	@echo "→ go vet"
	go vet $(PKGS)

test:
	@echo "→ go test"
	go test -timeout=30s $(PKGS)

test-verbose:
	@echo "→ go test -v"
	go test -v -timeout=30s $(PKGS)

test-single:
	@echo "→ go test $(TESTARGS)"
	go test -timeout=30s $(PKGS) $(TESTARGS)

coverage: $(COVERAGE_FILE)
	@echo "→ cobertura total:"
	go tool cover -func=$(COVERAGE_FILE)

$(COVERAGE_FILE):
	@echo "→ generando cobertura..."
	go test $(PKGS) -coverprofile=$(COVERAGE_FILE)

clean:
	@echo "→ limpiando..."
	rm -f $(COVERAGE_FILE)

