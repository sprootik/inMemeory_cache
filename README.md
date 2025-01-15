Eng | [Rus](./README-ru.md)   

This project implements a thread-safe in-memory lru cache with delete based on the element's lifetime.
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
goos: linux
goarch: amd64
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache/Add_element_in_cache-16           7528809               143.0 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8115110               142.3 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8180167               143.2 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8265849               143.2 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8229319               143.5 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16        60601836                18.59 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        226194962                5.226 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        222484065                5.221 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        228948388                5.284 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        226915117                5.258 ns/op           0 B/op          0 allocs/op
```