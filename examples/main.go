package main

import (
	"fmt"
	"sync"
	"time"

	cache "github.com/sprootik/inMemeory_cache"
)

func main() {
	// initiate a new cache with an item lifetime (TLRU) of 1 seconds. Key must be comparable
	// If desired, you can add time-to-live jitter to spread the load on the storage during cache eviction.
	c := cache.NewCache[int, string](1000).WithTimeout(1 * time.Second).WithJitter(time.Millisecond)

	// add 1001 elements in cache.
	// The cache will only have the last 1000 elements since the capacity is 1000
	var wg sync.WaitGroup
	for i := range 1001 {
		wg.Add(1)
		go func() {
			key := i
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
	for range 2 {
		key := 666
		value, ok := c.Get(key)
		if !ok {
			fmt.Printf("key %d not found\n", key)
		} else {
			fmt.Printf("key: %d found in cache, value: %s\n", key, value)
		}
		fmt.Printf("c size: %d\n", c.CacheSize())
		time.Sleep(1 * time.Second)
	}

	/* if you want to override ttl for an element you can do it when receiving using modifiers
	This will override the lifetime (without jitter) of the element when received,
	all preemptive logic will work based on the newTottle for this request
	*/

	for range 2 {
		key := 777
		value, ok := c.Get(key, cache.ChangeTTL(time.Hour))
		if !ok {
			fmt.Printf("key with change ttl %d not found\n", key)
		} else {
			fmt.Printf("key: %d with change ttl found in cache, value: %s\n", key, value)
		}
	}
	// show cache size
	fmt.Printf("c size: %d\n", c.CacheSize())
}
