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
BenchmarkCache/Add_element_in_cache-16           6219906               197.8 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6098492               200.0 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6096046               197.1 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6122575               197.7 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           5955610               196.8 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Get_element_from_cache-16        29190554                42.93 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        38835883                26.57 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        44114065                26.80 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        44893519                26.81 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        45058827                27.81 ns/op            7 B/op          0 allocs/op
```