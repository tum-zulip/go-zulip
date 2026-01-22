# Benchmarking go-zulip

This document describes the benchmark suite for go-zulip and how to use it effectively.

## Overview

The benchmark suite measures client-side performance of go-zulip using mock HTTP servers. Benchmarks focus on:

1. **Event Queue Operations** (`zulip/api/real_time_events`) - Most critical for bot performance
2. **Regular API Operations** (`zulip/client`) - Message sending, channel/user fetching

## Quick Start

```bash
# Run specific benchmarks (recommended)
go test -bench=BenchmarkRegisterQueue -benchmem ./zulip/api/real_time_events
go test -bench=BenchmarkSendMessage/Sequential -benchmem ./zulip/client

# Run all sequential benchmarks (safe)
go test -bench='Sequential' -benchmem ./zulip/client ./zulip/api/real_time_events

# Run with profiling
go test -bench=BenchmarkRegisterQueue -benchmem -cpuprofile=cpu.prof ./zulip/api/real_time_events
go tool pprof -http=:8080 cpu.prof
```

## Available Makefile Targets

```bash
make bench          # Run all benchmarks (may hit OS limits)
make bench-events   # Event queue benchmarks only
make bench-client   # Client API benchmarks only
make bench-profile  # Run with CPU/memory profiling
make bench-compare  # Compare with main branch using benchstat
```

## Benchmark Categories

### Event Queue Benchmarks (Priority)

Located in `zulip/api/real_time_events/api_real_time_events_bench_test.go`:

- **BenchmarkRegisterQueue** - One-time queue registration cost
- **BenchmarkGetEvents** - Event polling (most frequent bot operation)
  - Sequential
  - Concurrent-10 (10 goroutines)
  - Concurrent-25 (25 goroutines)
- **BenchmarkEventQueueListener** - Realistic bot polling loop
  - SingleConsumer
  - Concurrent-5/10 (multiple event consumers)
- **BenchmarkLongLivedQueue** - Many polls on single queue
- **BenchmarkMultipleQueues** - Multiple bot instances
- **BenchmarkEventProcessing** - Event unmarshaling overhead
- **BenchmarkDeleteQueue** - Queue cleanup

### Client API Benchmarks

Located in `zulip/client/client_bench_test.go`:

- **BenchmarkSendMessage** - Message sending (bot responses)
- **BenchmarkGetChannels** - Channel list fetching
- **BenchmarkGetUsers** - User list fetching
- **BenchmarkMixedWorkload** - Realistic bot workload (60% messages, 20% channels, 20% users)
- **BenchmarkClientInitialization** - Client creation cost

Each includes sequential and concurrent (10/25 goroutines) variants.

## Known Limitations

### Port Exhaustion on macOS

Running all concurrent benchmarks together may exhaust the OS ephemeral port range, causing `can't assign requested address` errors. This is a testing artifact, not a library issue.

**Workarounds:**

1. **Run benchmarks selectively** (recommended):
   ```bash
   # One benchmark at a time
   go test -bench=BenchmarkRegisterQueue -benchmem ./zulip/api/real_time_events
   go test -bench=BenchmarkSendMessage/Sequential -benchmem ./zulip/client
   ```

2. **Run only sequential benchmarks**:
   ```bash
   go test -bench='Sequential' -benchmem ./...
   ```

3. **Reduce benchmark time** for quick validation:
   ```bash
   go test -bench=. -benchtime=100x -benchmem ./zulip/client
   ```

4. **Run on Linux** - Linux has better ephemeral port management

### Why This Happens

- Each sub-benchmark creates a new mock HTTP server
- Concurrent benchmarks create many simultaneous connections
- macOS defaults to 16K-64K ephemeral ports with 15-second TIME_WAIT
- Sequential execution of many concurrent benchmarks can exhaust this range

This does **not** affect production usage - real bots use persistent connections to a single server.

## Interpreting Results

### Sample Output

```
BenchmarkRegisterQueue-8         17847    73307 ns/op    14208 B/op    140 allocs/op
BenchmarkGetEvents/Sequential-8  16539    71083 ns/op    17147 B/op    196 allocs/op
```

- **73307 ns/op** - 73 microseconds per operation (throughput: ~13,600 ops/sec)
- **14208 B/op** - 14KB allocated per operation
- **140 allocs/op** - 140 memory allocations per operation

### What to Look For

1. **High ns/op** - Slow operations that need optimization
2. **High B/op** - Memory-intensive operations (GC pressure)
3. **High allocs/op** - Allocation-heavy paths (escape analysis opportunities)
4. **Regressions** - Use `make bench-compare` to detect performance changes

### Typical Bot Performance Profile

From benchmark results, typical performance characteristics:

- **Register queue**: ~70-80µs, ~14KB, ~140 allocs
- **Poll events (empty)**: ~70-75µs, ~17KB, ~196 allocs
- **Send message**: ~65-70µs, ~15KB, ~178 allocs
- **Get channels**: ~90-95µs, ~42KB, ~216 allocs (larger response)
- **Get users**: ~75-80µs, ~21KB, ~166 allocs

These are client-side costs only (mock server has zero latency). Real-world performance includes network RTT + server processing.

## Best Practices

### For Development

1. **Run specific benchmarks** during development:
   ```bash
   go test -bench=BenchmarkSendMessage/Sequential -benchmem ./zulip/client
   ```

2. **Use profiling** for optimization work:
   ```bash
   make bench-profile
   # Opens pprof web UI on :8080
   ```

3. **Compare before/after** changes:
   ```bash
   # Before changes
   go test -bench=BenchmarkGetEvents/Sequential -benchmem ./zulip/api/real_time_events > old.txt

   # After changes
   go test -bench=BenchmarkGetEvents/Sequential -benchmem ./zulip/api/real_time_events > new.txt

   # Compare
   benchstat old.txt new.txt
   ```

### For CI/CD

- Run sequential benchmarks only to avoid flakiness
- Use `-benchtime=100x` for faster CI runs
- Store results as artifacts for historical comparison
- Set up performance regression alerts using benchstat

## Example Workflow

```bash
# 1. Establish baseline
go test -bench=BenchmarkSendMessage/Sequential -benchmem ./zulip/client > baseline.txt

# 2. Make changes to the codebase
# ... edit files ...

# 3. Re-run benchmark
go test -bench=BenchmarkSendMessage/Sequential -benchmem ./zulip/client > new.txt

# 4. Compare
benchstat baseline.txt new.txt

# Example output:
# name                          old time/op  new time/op  delta
# SendMessage/Sequential-8      66.6µs ± 2%  62.1µs ± 1%  -6.76%
#
# name                          old alloc/op new alloc/op delta
# SendMessage/Sequential-8      14.9kB ± 0%  14.2kB ± 0%  -4.70%
```

## Extending the Suite

To add new benchmarks:

1. Follow existing patterns in test files
2. Use `setupEventMockServer()` or `setupAPIMockServer()` for mock servers
3. Use `createBenchmarkClient()` for client creation
4. Include both sequential and concurrent variants
5. Add realistic sub-benchmarks rather than synthetic micro-benchmarks

Example:
```go
func BenchmarkNewOperation(b *testing.B) {
    server := setupAPIMockServer()
    defer server.Close()

    c := createBenchmarkClient(b, server.URL)
    ctx := context.Background()

    b.Run("Sequential", func(b *testing.B) {
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            _, _, err := c.NewOperation(ctx).Execute()
            if err != nil {
                b.Fatalf("NewOperation failed: %v", err)
            }
        }
    })
}
```

## References

- [Go benchmark documentation](https://pkg.go.dev/testing#hdr-Benchmarks)
- [benchstat tool](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)
- [pprof profiling](https://github.com/google/pprof/blob/main/doc/README.md)
