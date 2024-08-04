package main

import (
	"fmt"
	"github.com/sprootik/inMemeory_cache"
	"sync"
	"time"
)

func main() {
	// initiate a new cache with an item lifetime of 1 seconds
	c := cache.NewCache(1000, 1*time.Second)

	// add 1001 elements in cache.
	// The cache will only have the last 1000 elements since the capacity is 1000
	var wg sync.WaitGroup
	for i := 0; i < 1001; i++ {
		wg.Add(1)
		go func() {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			c.Add(key, value)
			wg.Done()
		}()
	}
	wg.Wait()

	// show cache size
	fmt.Printf("c size: %d\n", c.CacheSize())

	// get value from cache by key. After second.
	// after a second, the elements are considered not relevant, since the lifetime is 1 second.
	// The element will be removed from the cache upon request
	for i := 0; i < 2; i++ {
		key := "key-666"
		value, ok := c.Get(key)
		if !ok {
			fmt.Println("key not found")
		} else {
			fmt.Printf("key: %s found in cache, value: %s\n", key, value)
		}
		fmt.Printf("c size: %d\n", c.CacheSize())
		time.Sleep(1 * time.Second)
	}
}
