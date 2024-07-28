Eng | [Rus](./README-ru.md)   

This project implements a thread-safe in-memory cache with delete based on the element's lifetime.
Algorithmic complexity of work is O(1). 

*** 
### Installation  
```
go get github.com/sprootik/inMemeory_cache@latest
```

***
### Examples of using  
See [example](./examples/main.go)
***
### Performance test    

```
=== CPU bench ===
goos: linux
goarch: amd64
pkg: cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache
BenchmarkCache-16         367285              3603 ns/op
PASS
ok      cache   1.379s
==================
=== MEM bench ===
goos: linux
goarch: amd64
pkg: cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache-16         314948              3264 ns/op             671 B/op          8 allocs/op
PASS
ok      cache   1.087s
```