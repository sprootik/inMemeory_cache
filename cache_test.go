package cache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

var cache *Cache

func init() {
	cache = NewCache(&sync.RWMutex{}, 1*time.Second)
	err := cache.StartCacheExecutor()
	if err != nil {
		panic(err)
	}
}

func printCacheElement() {
	cache.mu.RLock()
	fmt.Printf("Cache tail: %p, head: %p\n", cache.tail, cache.head)

	fmt.Println("----------")
	for _, v := range cache.index {
		fmt.Printf("Pointer: %p\n Node: %+v\n", v, v)
	}
	cache.mu.RUnlock()
	fmt.Println("----------")
}

func TestCacheCorrectWork(t *testing.T) {
	cache.mu.Lock()
	cache.lifeTime = 6 * time.Second
	cache.mu.Unlock()

	for i := 0; i < 5; i++ {
		func() {
			key := fmt.Sprintf("%d", i)
			value := fmt.Sprintf("value-%d", i)
			cache.Add(key, value)
			fmt.Printf("add Node: k:%s v:%s\n", key, value)
			time.Sleep(1 * time.Second)
		}()
	}

	fmt.Printf("Cache size: %d\n", cache.CacheSize())
	printCacheElement()

	element, ok := cache.Get(fmt.Sprintf("%d", 4))
	fmt.Println("**********")
	if ok {
		fmt.Printf("element in cache: %s\n", element)
	} else {
		fmt.Println("not found in cache")
	}
	fmt.Println("**********")

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)

		fmt.Printf("Cache size: %d\n", cache.CacheSize())
		printCacheElement()
	}
}

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			cache.Add(key, value)
			cache.Get(key)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Cache size: %d\n", cache.CacheSize())
	time.Sleep(1 * time.Second)
	fmt.Printf("Cache size: %d\n", cache.CacheSize())

}

func BenchmarkCache(b *testing.B) {
	var wg sync.WaitGroup
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		wg.Add(1)
		go func() {
			key := strconv.Itoa(j)
			cache.Add(key, "value")
			cache.Get(key)
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}
