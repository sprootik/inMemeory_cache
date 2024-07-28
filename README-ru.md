Данный проект реализует потокобезопасный in-memory кэш, с отчисткой по времени жизни элемента.  
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
pkg: cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache
BenchmarkCache-16         367285              3603 ns/op
PASS
ok      cache   1.379s
==================
=== MEM bench ===
goos: linux
goarch: amd64
pkg: cache
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkCache-16         314948              3264 ns/op             671 B/op          8 allocs/op
PASS
ok      cache   1.087s
```
