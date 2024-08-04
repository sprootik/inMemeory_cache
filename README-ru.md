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
BenchmarkCache/Add_element_in_cache-16           6152816               195.5 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6090147               196.3 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6015614               196.1 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6138055               196.4 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Add_element_in_cache-16           6079730               197.4 ns/op            89 B/op          2 allocs/op
BenchmarkCache/Get_element_from_cache-16        29385784                44.97 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        45410721                26.71 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        44493037                26.47 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        45292598                30.81 ns/op            7 B/op          0 allocs/op
BenchmarkCache/Get_element_from_cache-16        43584282                26.34 ns/op            7 B/op          0 allocs/op
```
