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
BenchmarkCache/Add_element_in_cache-16           1553043               682.1 ns/op           235 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1834908               735.9 ns/op           162 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1791330               742.9 ns/op           145 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1768822               680.7 ns/op           146 B/op          3 allocs/op
BenchmarkCache/Add_element_in_cache-16           1695360               684.6 ns/op           146 B/op          3 allocs/op
BenchmarkCache/Get_element_from_cache-16         7555476               163.5 ns/op            31 B/op          2 allocs/op
BenchmarkCache/Get_element_from_cache-16         7170780               166.4 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7251628               155.3 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7252480               156.6 ns/op            31 B/op          1 allocs/op
BenchmarkCache/Get_element_from_cache-16         7300053               157.1 ns/op            31 B/op          1 allocs/op
```
