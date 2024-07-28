// Package cache
/*
the package implements in-memory cache with element life time.
This cache is thread safe.
Algorithmic complexity of work is O(1).
*/
package cache

import (
	"fmt"
	"sync"
	"time"
)

var (
	once         sync.Once
	instantiated *Cache
)

// node linked list cache element
type node struct {
	key      string
	value    interface{}
	time     time.Time
	next     *node
	previous *node
}

// Cache cache structure
type Cache struct {
	size          int
	head          *node
	tail          *node
	index         map[string]*node
	mu            *sync.RWMutex
	lifeTime      time.Duration
	isRunExecutor bool
}

// NewCache init new cache, it's the singleton
func NewCache(mu *sync.RWMutex, lifeTime time.Duration) *Cache {
	once.Do(func() {
		instantiated = &Cache{
			mu: mu,
			//head, tail:  nil,
			//size:  0,
			index:    make(map[string]*node),
			lifeTime: lifeTime,
		}
	})
	return instantiated
}

// IsRunExecutor is the clear cache item function running?
func (c *Cache) IsRunExecutor() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.isRunExecutor
}

// CacheSize number of items in the cache at the moment
func (c *Cache) CacheSize() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.size
}

// Add add the element in cache
func (c *Cache) Add(key string, value interface{}) {
	start := time.Now()
	element := &node{key: key, value: value, time: start}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.head == nil && c.tail == nil {
		c.head = element
		c.tail = element
	} else {
		element.previous = c.head
		c.head.next = element
		c.head = element
	}

	c.index[key] = element
	c.size++
}

// Get get from cache by key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	val, ok := c.index[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return val.value, true
}

func (c *Cache) cacheExecutor() {
	c.mu.RLock()
	cacheSize := c.size
	tailElement := c.tail
	lifeTime := c.lifeTime
	c.mu.RUnlock()

	if cacheSize > 0 {
		if time.Since(tailElement.time) > lifeTime {
			c.mu.Lock()
			defer c.mu.Unlock()

			delete(c.index, c.tail.key)
			c.size--
			if c.size == 0 {
				c.head = nil
				c.tail = nil
				return
			}
			c.tail = c.tail.next
			c.tail.previous = nil
		}
	}
}

// StartCacheExecutor launch the function to clear cache items
func (c *Cache) StartCacheExecutor() error {
	isRun := c.IsRunExecutor()
	if isRun {
		return fmt.Errorf("cache executor is already running")
	}

	c.mu.Lock()
	c.isRunExecutor = true
	c.mu.Unlock()

	go func() {
		for {
			c.cacheExecutor()
		}
	}()
	return nil
}
