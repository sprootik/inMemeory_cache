testCacheRace:
	go test -count=1 --race -run ^TestCache$$ -v .

testCacheCPU:
	go test -count=1 -run ^TestCache$$ -cpuprofile tests/cpu.out
	go tool pprof -http=:8080 tests/cpu.out

testCacheMEM:
	go test -count=1 -run ^TestCache$$ -memprofile tests/mem.out
	go tool pprof -http=:8080 tests/mem.out

testCacheMutex:
	go test -count=1 -run ^TestCache$$ --mutexprofile tests/mutex.out
	go tool pprof -http=:8080 tests/mutex.out

testCacheTrace:
	go test -count=1 -run ^TestCache$$ --trace  tests/trace.out
	go tool trace tests/trace.out

testCacheBench:
	@go test -benchmem -count=5 -run=^$$ -bench ^BenchmarkCache$$ .