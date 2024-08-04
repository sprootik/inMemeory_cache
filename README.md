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
BenchmarkCache/Add_element_in_cache-16           1590344               760.7 ns/op           243 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1804701               714.7 ns/op           146 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1805014               737.6 ns/op           146 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1774330               735.4 ns/op           145 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1740439               746.2 ns/op           145 B/op          3 allocs/op
BenchmarkCache/Get_element_from_cache-16         7495243               161.5 ns/op            31 B/op          2 allocs/op
BenchmarkCache/Get_element_from_cache-16         7262822               155.8 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7327350               154.6 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7273894               155.6 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7252576               155.7 ns/op            31 B/op          1 allocs/op
```