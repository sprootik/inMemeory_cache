testCacheCorrect:
	go test -count=1 -run ^TestCorrectWork$$ -v .

testCacheRace:
	go test -count=1 --race -run ^TestCache$$ -v .

testCacheJitter:
	go test -count=1 -run ^TestJitter$$ -v .

testCacheBench:
	@go test -benchmem -count=5 -run=^$$ -bench ^BenchmarkCache$$ .

pprofCacheCPU:
	go test -count=1 -run ^TestCache$$ -cpuprofile pprof/cpu.out
	go tool pprof -http=:8080 pprof/cpu.out

pprofCacheMEM:
	go test -count=1 -run ^TestCache$$ -memprofile pprof/mem.out
	go tool pprof -http=:8080 pprof/mem.out

pprofCacheMutex:
	go test -count=1 -run ^TestCache$$ --mutexprofile pprof/mutex.out
	go tool pprof -http=:8080 pprof/mutex.out

pprofCacheTrace:
	go test -count=1 -run ^TestCache$$ --trace  pprof/trace.out
	go tool trace pprof/trace.out