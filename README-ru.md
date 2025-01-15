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
