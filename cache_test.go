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
	cache = NewCache(&sync.RWMutex{}, 1000, 1*time.Second)
}

func printCacheElement() {
	cache.mu.RLock()
	fmt.Printf("Cache tail: %p, head: %p\n", cache.tail, cache.head)

	fmt.Println("----------")
	fmt.Printf("data: %p %+v\n", cache.data, cache.data)
	for _, v := range cache.data {
		fmt.Printf("Pointer: %p\n Node: %+v\n", v, v)
	}
	cache.mu.RUnlock()
	fmt.Println("----------")
}

func TestCacheCorrectWork(t *testing.T) {
	cache.mu.Lock()
	cache.capacity = 4
	cache.mu.Unlock()

	for i := 0; i < 5; i++ {
		func() {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			cache.Add(key, value)
			fmt.Printf("add Node: k:%s v:%s\n", key, value)
		}()
	}

	fmt.Printf("Cache size: %d\n", cache.CacheSize())
	printCacheElement()

	for i := 0; i < 2; i++ {
		element, ok := cache.Get(fmt.Sprintf("key-%d", 3))
		fmt.Println("**********")
		if ok {
			fmt.Printf("element in cache: %s\n", element)
		} else {
			fmt.Println("not found in cache")
		}
		fmt.Println("**********")

		time.Sleep(1 * time.Second)
		fmt.Printf("Cache size: %d\n", cache.CacheSize())
		printCacheElement()
		time.Sleep(1 * time.Second)
	}
}

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	//add & get
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			key := strconv.Itoa(i)
			value := "value-" + key
			cache.Add(key, value)
			cache.Get(key)
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)

	//get & del
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			key := strconv.Itoa(i)
			cache.Get(key)
			wg.Done()
		}()
	}
	wg.Wait()

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
