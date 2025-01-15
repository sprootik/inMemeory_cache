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
BenchmarkCache/Add_element_in_cache-16           6118148               192.8 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           6124830               193.0 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           7820300               143.7 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8167287               142.7 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           8197624               146.5 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16        71912908                16.03 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        169807807                5.947 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        200084960                5.925 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        202126875                5.855 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        203759926                5.907 ns/op           0 B/op          0 allocs/op
```