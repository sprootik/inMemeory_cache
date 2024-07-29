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
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache
BenchmarkCache-16         975109              1543 ns/op
PASS
ok      github.com/sprootik/inMemeory_cache     1.537s
==================
=== MEM bench ===
goos: linux
goarch: amd64
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache-16        1000000              1544 ns/op             681 B/op          4 allocs/op
PASS
ok      github.com/sprootik/inMemeory_cache     1.578s
```