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
BenchmarkCache/Add_element-16            7307991               159.6 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6952244               166.7 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            7123704               166.9 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            7097658               167.5 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            7045665               162.8 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Get_element-16           51850240                22.77 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           50203240                22.80 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           52538589                22.78 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           50939607                22.78 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           52898586                22.78 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              168130641                6.256 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              191839611                6.261 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              191384476                6.262 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              191603686                6.266 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              191220027                6.270 ns/op           0 B/op          0 allocs/op
```