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

// node linked list cache element
type node struct {
	key      string
	value    any
	time     time.Time
	next     *node
	previous *node
}

// Cache cache structure
type Cache struct {
	size     int
	capacity int
	head     *node
	tail     *node
	data     map[string]*node
	mu       *sync.RWMutex
	lifeTime time.Duration
}

/*
NewCache init new cache.

	capacity: this is the capacity of the lru cache. If the number of added elements is greater than the capacity the last element is removed

	nodeLifeTime: cache item lifetime
*/
func NewCache(mu *sync.RWMutex, capacity int, nodeLifeTime time.Duration) *Cache {
	return &Cache{
		mu: mu,
		//head, tail:  nil,
		//size:  0,
		capacity: capacity,
		data:     make(map[string]*node, capacity),
		lifeTime: nodeLifeTime,
	}
}

// CacheSize number of items in the cache at the moment
func (c *Cache) CacheSize() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.size
}

// CacheCapacity current cache capacity
func (c *Cache) CacheCapacity() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.capacity
}

// SetCapacity set cache capacity
func (c *Cache) SetCapacity(capacity int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.capacity = capacity
}

// Add add the element in cache
func (c *Cache) Add(key string, value any) {
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
	c.size++
}

// Get get from cache by key
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	node, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if time.Since(node.time) > c.lifeTime {
		c.unsafeRemove(node)
		delete(c.data, key)
		return nil, false
	}

	return node.value, true
}

// unsafeRemove thread-unsafe removal of an element from a linked-list
func (c *Cache) unsafeRemove(node *node) {
	if node.next != nil && node.previous != nil {
		node.previous.next = node.next
		node.next.previous = node.previous
		delete(c.data, node.key)
	} else if node.previous == nil && node.next != nil {
		c.tail = node.next
		node.next.previous = nil
		delete(c.data, node.key)
	} else if node.next == nil && node.previous != nil {
		c.head = node.previous
		node.previous.next = nil
		delete(c.data, node.key)
	} else {
		delete(c.data, node.key)
		c.head = nil
		c.tail = nil
	}
	c.size--
}
