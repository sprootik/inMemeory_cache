Данный проект реализует потокобезопасный in-memory кэш, с отчисткой по времени.  
Алгоритмическая сложностью работы O(1).

*** 
### Установка

***
### Примеры использования
***
### Тест производительности  

#### CPU 

```
added 360772 elements in cache
Cache size befor clean: 360348, Cache size after clean: 0
  360772              3897 ns/op
```
#### Memory
```
added 352123 elements in cache
Cache size befor clean: 351749, Cache size after clean: 0
  352123              3664 ns/op             667 B/op          8 allocs/op
```
