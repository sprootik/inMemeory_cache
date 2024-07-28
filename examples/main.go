package main

import (
	"fmt"
	"github.com/sprootik/inMemeory_cache"
	"sync"
	"time"
)

func main() {
	// initiate a new cache with an item lifetime of 1 seconds
	c := cache.NewCache(&sync.RWMutex{}, 1*time.Second)

	// start cache executor
	err := c.StartCacheExecutor()
	if err != nil {
		panic(err)
	}

	// add 1000 elements in cache
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			c.Add(key, value)
			c.Get(key)
			wg.Done()
		}()
	}
	wg.Wait()

	// show cache size
	fmt.Printf("c size: %d\n", c.CacheSize())

	// get value from cache by key
	key := "key-666"
	value, ok := c.Get(key)
	if !ok {
		fmt.Println("key not found")
	} else {
		fmt.Printf("key: %s found in cache, value: %s\n", key, value)
	}

	// waiting for c to be cleared
	time.Sleep(1 * time.Second)
	fmt.Printf("c size: %d\n", c.CacheSize())
}
