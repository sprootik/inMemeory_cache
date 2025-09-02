Данный проект реализует простой потокобезопасный in-memory LRU/TLRU кэш,
Алгоритмическая сложностью работы O(1).

*** 
### Установка  
```
go get github.com/sprootik/inMemeory_cache@latest
```

***
### Документация
```
https://pkg.go.dev/github.com/sprootik/inMemeory_cache
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
BenchmarkCache/Add_element-16           10140134               116.6 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9573124               124.0 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9321261               124.2 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            9260470               123.8 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Add_element-16            7964702               171.4 ns/op            32 B/op          1 allocs/op
BenchmarkCache/Get_element-16           37832053                31.65 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           38113810                31.57 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           48150932                22.21 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           50007607                22.40 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element-16           51462730                22.34 ns/op            0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              229280216                5.242 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              228699200                5.444 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              223116471                5.355 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              218576280                5.374 ns/op           0 B/op          0 allocs/op
BenchmarkCache/Get_element_with_timeout-16              221789240                5.397 ns/op           0 B/op          0 allocs/op
```
