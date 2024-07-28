Eng | [Rus](./README-ru.md)   

This project implements a thread-safe in-memory thread with time-based cache cleanup.
Algorithmic classification of O(1) work.

*** 
### Installation  

***
### Examples of using  
***
### Performance test    

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