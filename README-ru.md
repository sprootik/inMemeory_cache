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
=== CPU bench ===
goos: linux
goarch: amd64
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache
BenchmarkCache-16        1000000              1154 ns/op
PASS
ok      github.com/sprootik/inMemeory_cache     1.231s
==================
=== MEM bench ===
goos: linux
goarch: amd64
pkg: github.com/sprootik/inMemeory_cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache-16        1000000              1163 ns/op             442 B/op          4 allocs/op
PASS
ok      github.com/sprootik/inMemeory_cache     1.227s
```
