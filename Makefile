testCacheRace:
	go test -count=1 --race -run ^TestCache$$ -v .

testCacheCPU:
	go test -count=1 -run ^TestCache$$ -cpuprofile tests/cpu.out
	go tool pprof -http=:8080 tests/cpu.out

testCacheMutex:
	go test -count=1 -run ^TestCache$$ --mutexprofile tests/mutex.out
	go tool pprof -http=:8080 tests/mutex.out

testCacheTrace:
	go test -count=1 -run ^TestCache$$ --trace  tests/trace.out
	go tool trace tests/trace.out

testCacheBench:
	@echo "=== CPU bench ==="
	@go test -v -bench=. -run ^BenchmarkCache$$ .
	@echo "=================="
	@echo "=== MEM bench ==="
	@go test -benchmem -run=^$$ -bench ^BenchmarkCache$$ .