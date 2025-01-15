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
BenchmarkCache/Add_element-16            6135909               192.6 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6195796               193.0 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6178669               193.1 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6179636               192.5 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6149724               192.2 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Get_element-16           232902384                5.204 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           221614870                5.166 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           232680560                5.176 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           226907667                5.435 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           227015454                5.665 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              224276396                5.667 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              222030883                5.373 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              221273929                5.333 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              221545132                6.254 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              167375360                7.172 ns/op           0 B/op          0 allocs/op
```