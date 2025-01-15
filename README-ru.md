Данный проект реализует потокобезопасный in-memory вытесняющий кэш, с отчисткой по времени жизни элемента.  
Алгоритмическая сложностью работы O(1).

*** 
### Установка  
```
go get github.com/sprootik/inMemeory_cache@latest
```
***
### Примеры использования
Смотри [example](./examples/main.go)
***
### Тест производительности  

```
goos: linux
goarch: amd64
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache/Add_element_in_cache-16           7344130               148.6 ns/op            64 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           7923462               158.2 ns/op            64 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           7925038               160.9 ns/op            64 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           5910595               199.9 ns/op            64 B/op          1 allocs/op
BenchmarkCache/Add_element_in_cache-16           5963118               198.9 ns/op            64 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16        51909685                22.24 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        155644192                7.635 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        161705904                7.400 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        162728091                7.483 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        163964517                7.361 ns/op           0 B/op          0 allocs/op
```
