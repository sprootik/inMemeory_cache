// Package cache
/*
the package implements in-memory LRU cache consisting of elements with life time.
This cache is thread safe.
Algorithmic complexity of work is O(1).
*/
package cache

import (
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
	capacity      int
	head          *node
	tail          *node
	data          map[string]*node
	mu            *sync.RWMutex
	lifeTime      time.Duration
	isRunExecutor bool
}

// NewCache init new cache, it's the singleton
func NewCache(mu *sync.RWMutex, capacity int, lifeTime time.Duration) *Cache {
	once.Do(func() {
		instantiated = &Cache{
			mu: mu,
			//head, tail:  nil,
			//size:  0,
			capacity: capacity,
			data:     make(map[string]*node),
			//index:    make(map[int]*node),
			lifeTime: lifeTime,
		}
	})
	return instantiated
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

	// delete of the latter when the size is exceeded
	if c.size >= c.capacity {
		c.unsafeRemove(c.tail)
	}

	if _, ok := c.data[key]; ok {
		return
	}

	if c.head == nil && c.tail == nil {
		c.head = element
		c.tail = element
	} else {
		element.previous = c.head
		c.head.next = element
		c.head = element
	}

	c.data[key] = element
	//c.index[c.size] = element
	c.size++
}

// Get get from cache by key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	val, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return val.value, true
}

func (c *Cache) unsafeRemove(node *node) {
	if node.next == nil && node.previous == nil {
		node.previous.next = node.next.previous
		node.next.previous = node.previous.next
		delete(c.data, node.key)
	} else if node.previous == nil && node.next != nil {
		node.next.previous = nil
		delete(c.data, node.key)
	} else if node.next == nil && node.previous != nil {
		node.previous.next = nil
		delete(c.data, node.key)
	}
	c.size--
}

//
//func (c *Cache) cacheExecutor() {
//	c.mu.RLock()
//	cacheSize := c.size
//	tailElement := c.tail
//	lifeTime := c.lifeTime
//	c.mu.RUnlock()
//
//	if cacheSize > 0 {
//		if time.Since(tailElement.time) > lifeTime {
//			c.mu.Lock()
//			defer c.mu.Unlock()
//
//			delete(c.data, c.tail.key)
//			c.size--
//			if c.size == 0 {
//				c.head = nil
//				c.tail = nil
//				return
//			}
//			c.tail = c.tail.next
//			c.tail.previous = nil
//		}
//	}
//}
