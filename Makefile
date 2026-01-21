.PHONY: test lint fmt coverage clean build install tidy check all
.PHONY: dev-server-start dev-server-provision dev-server-run dev-server-status dev-server-stop

COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html
DEV_SERVER_SCRIPT := ./scripts/run-dev-zerver.sh
DEV_SERVER_REF := main
DEV_SERVER_DIR := /tmp/zulip-zerver/$(DEV_SERVER_REF)/zulip

test:
	go test ./... -v -timeout 10m

lint:
	golangci-lint run

fmt:
	golangci-lint run --fix

coverage:
	go test ./... -covermode=atomic -coverprofile=$(COVERAGE_OUT) -coverpkg=./... -timeout 10m
	go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

build:
	go build ./...

install:
	go mod download

tidy:
	go mod tidy
	go mod verify

clean:
	go clean
	rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)

dev-server-start:
	sudo $(DEV_SERVER_SCRIPT) --ref $(DEV_SERVER_REF)

dev-server-provision:
	sudo $(DEV_SERVER_SCRIPT) --ref $(DEV_SERVER_REF) --provision-only

dev-server-run:
	$(DEV_SERVER_SCRIPT) --ref $(DEV_SERVER_REF) --run-only

dev-server-status:
	@pgrep -af "run-dev" || echo "Not running"

dev-server-stop:
	@pkill -f "run-dev" || echo "Not running"

check: lint test

all: tidy fmt lint test build
