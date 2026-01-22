.PHONY: test lint fmt coverage clean build install tidy check all
.PHONY: dev-server-start dev-server-provision dev-server-run dev-server-status dev-server-stop
.PHONY: bench bench-events bench-client bench-profile bench-compare

COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html
DEV_SERVER_SCRIPT := ./scripts/run-dev-zerver.sh
DEV_SERVER_REF := main
DEV_SERVER_DIR := /tmp/zulip-zerver/$(DEV_SERVER_REF)/zulip

test:
	go test ./... -v -timeout 10m

bench:
	@echo "Running all benchmarks..."
	go test -bench=. -benchmem -run=^$$ ./zulip/client ./zulip/api/real_time_events

bench-events:
	@echo "Running event queue benchmarks..."
	go test -bench=. -benchmem -run=^$$ ./zulip/api/real_time_events

bench-client:
	@echo "Running client API benchmarks..."
	go test -bench=. -benchmem -run=^$$ ./zulip/client

bench-profile:
	@echo "Running benchmarks with CPU and memory profiling..."
	@mkdir -p benchmarks
	go test -bench=. -benchmem -cpuprofile=benchmarks/cpu.prof -memprofile=benchmarks/mem.prof -run=^$$ ./zulip/client ./zulip/api/real_time_events
	@echo ""
	@echo "Profiles saved to benchmarks/ directory"
	@echo "View CPU profile:    go tool pprof -http=:8080 benchmarks/cpu.prof"
	@echo "View memory profile: go tool pprof -http=:8080 benchmarks/mem.prof"

bench-compare:
	@echo "Comparing benchmarks with main branch..."
	@which benchstat > /dev/null || (echo "Installing benchstat..." && go install golang.org/x/perf/cmd/benchstat@latest)
	@mkdir -p benchmarks
	@echo "Running current benchmarks..."
	@go test -bench=. -benchmem -run=^$$ ./zulip/client ./zulip/api/real_time_events > benchmarks/current.txt
	@echo "Stashing changes and running main branch benchmarks..."
	@git stash -q && go test -bench=. -benchmem -run=^$$ ./zulip/client ./zulip/api/real_time_events > benchmarks/main.txt && git stash pop -q || (git stash pop -q && exit 1)
	@echo ""
	@echo "Benchmark comparison (main vs current):"
	@echo "========================================"
	@benchstat benchmarks/main.txt benchmarks/current.txt

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
