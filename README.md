Eng | [Rus](./README-ru.md)   

This project implements a simple thread-safe in-memory LRU/TLRU cache.
Algorithmic complexity of work is O(1). 

*** 
### Installation  
```
go get github.com/sprootik/inMemeory_cache@latest
```

***
### Documentation
```
https://pkg.go.dev/github.com/sprootik/inMemeory_cache
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
BenchmarkCache/Add_element-16           10018486               116.2 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9467871               125.0 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9155750               126.1 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9351379               127.1 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9317839               126.2 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Get_element-16           51341473                23.12 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           49172379                23.17 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           50471332                23.20 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           50269903                23.08 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           49998426                23.35 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              162310446                6.395 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              229031422                5.224 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              228467642                5.250 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              225184285                5.289 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              219739296                5.314 ns/op           0 B/op          0 allocs/op
```