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
BenchmarkCache/Add_element-16            6070723               195.1 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6056997               195.0 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            6042778               189.6 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            7034952               144.2 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Add_element-16            8045034               144.6 ns/op            48 B/op          1 allocs/op
BenchmarkCache/Get_element-16           77292549                15.40 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           186223486                5.440 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           220376948                5.428 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           222435225                5.471 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           218950906                5.359 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              225407926                5.401 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              220393656                5.371 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              222802556                5.447 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              219089589                6.745 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              162776655                7.412 ns/op           0 B/op          0 allocs/op
```
