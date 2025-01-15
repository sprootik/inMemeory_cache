package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cache *Cache[int, int]

func init() {
	cache = NewCache[int, int](1000, 1*time.Second)
}

// func printCacheElement() {
// 	cache.mu.Lock()
// 	fmt.Printf("Cache tail: %p, head: %p\n", cache.tail, cache.head)

// 	fmt.Println("----------")
// 	fmt.Printf("data: %p %+v\n", cache.data, cache.data)
// 	for _, v := range cache.data {
// 		fmt.Printf("Pointer: %p\n Node: %+v\n", v, v)
// 	}
// 	cache.mu.Unlock()
// 	fmt.Println("----------")
// }

// func TestCacheCorrectWork(t *testing.T) {
// 	cache.mu.Lock()
// 	cache.capacity = 4
// 	cache.mu.Unlock()

// 	for i := 0; i < 5; i++ {
// 		func() {
// 			key := fmt.Sprintf("key-%d", i)
// 			value := fmt.Sprintf("value-%d", i)
// 			cache.Add(key, value)
// 			fmt.Printf("*add Node: k:%s v:%s\n", key, value)
// 		}()
// 	}

// 	fmt.Printf("Cache size: %d\n", cache.CacheSize())
// 	printCacheElement()

// 	for i := 0; i < 2; i++ {
// 		fmt.Printf("*get iteration, second: %d\n", i)
// 		element, ok := cache.Get(fmt.Sprintf("key-%d", 3))
// 		fmt.Println("**********")
// 		if ok {
// 			fmt.Printf("element in cache: %s\n", element)
// 		} else {
// 			fmt.Println("not found in cache")
// 		}
// 		fmt.Println("**********")

// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("Cache size: %d\n", cache.CacheSize())
// 		printCacheElement()
// 		time.Sleep(1 * time.Second)
// 	}

// 	key := fmt.Sprintf("key-%d", 2)
// 	value := fmt.Sprintf("value-%d", 2)
// 	cache.Add(key, value)
// 	fmt.Printf("*add Node after Get: k:%s v:%s\n", key, value)
// 	fmt.Printf("Cache size: %d\n", cache.CacheSize())
// 	printCacheElement()
// }

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	//add & get
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			cache.Add(i, i)
			cache.Get(i)
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)

	//get & del
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			cache.Get(i)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Cache size: %d\n", cache.CacheSize())
}

func BenchmarkCache(b *testing.B) {
	b.ResetTimer()
	b.Run("Add element in cache", func(b *testing.B) {

		for j := 0; j < b.N; j++ {
			func() {
				cache.Add(j, j)
			}()
		}
	})

	b.Run("Get element from cache", func(b *testing.B) {

		for j := 0; j < b.N; j++ {
			func() {
				cache.Get(j)
			}()
		}
	})
}

func printCache() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	fmt.Printf("head: %p, tail %p\n", cache.head, cache.tail)
	h := cache.head // change hed & tail
	var i int
	for h != nil {
		fmt.Printf("%d %p %+v\n", i, h, h)
		h = h.next
		i++
	}
	fmt.Println("=======")
}
func TestWork(t *testing.T) {
	printCache()
	cache.Add(0, 0)
	cache.Add(1, 1)
	cache.Add(2, 2)
	cache.Add(3, 3)
	printCache()
	cache.Get(1)
	printCache()
	printCache()
	cache.Get(3)
	printCache()
}
