testCacheRace:
	go test -count=1 --race -run ^TestCache$$ -v ./cache

testCacheCPU:
	cd cache; \
	go test -count=1 -run ^TestCache$$ -cpuprofile tests/cpu.out
	go tool pprof -http=:8080 cache/tests/cpu.out

testCacheMutex:
	cd pkg/cache; \
	go test -count=1 -run ^TestCache$$ --mutexprofile tests/mutex.out
	go tool pprof -http=:8080 cache/tests/mutex.out

testCacheTrace:
	cd pkg/cache; \
	go test -count=1 -run ^TestCache$$ --trace  tests/trace.out
	go tool trace cache/tests/trace.out

testCacheBench:
	@echo "=== CPU benc ==="
	@go test -bench=. -run ^BenchmarkCache$$ cache/cache
	@echo "================++"
	@echo "=== MEM bench ==="
	@go test -benchmem -run=^$$ -bench ^BenchmarkCache$$ cache/cache