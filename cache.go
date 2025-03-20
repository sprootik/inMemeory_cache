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
type node[K comparable, V any] struct {
	key      K
	value    V
	time     time.Time
	next     *node[K, V]
	previous *node[K, V]
}

// Cache cache structure
type Cache[K comparable, V any] struct {
	capacity int
	head     *node[K, V]
	tail     *node[K, V]
	data     map[K]*node[K, V]
	mu       sync.Mutex
	lifeTime time.Duration
	zeroVal  V
}

// unsafeAddToTail thread-unsafe add element to end of linked-list
func (c *Cache[K, V]) unsafeAddToTail(node *node[K, V]) {
	c.data[node.key] = node

	if c.head == nil && c.tail == nil {
		c.head = node
		c.tail = node
		return
	}
	node.previous = c.tail
	c.tail.next = node
	c.tail = node
}

// unsafeDelete thread-unsafe removal of an element from a linked-list
func (c *Cache[K, V]) unsafeDelete(node *node[K, V]) {
	// [] - [x] - []
	if node.next != nil && node.previous != nil {
		node.previous.next = node.next
		node.next.previous = node.previous
		// [x] - [] - []
	} else if node.previous == nil && node.next != nil {
		c.head = node.next
		node.next.previous = nil
		// [] - [] - [x]
	} else if node.next == nil && node.previous != nil {
		c.tail = node.previous
		node.previous.next = nil
	} else {
		c.head = nil
		c.tail = nil
	}
	delete(c.data, node.key)
}

// unsafeMoveToTail thread-unsafely move element to end of linked list
func (c *Cache[K, V]) unsafeMoveToTail(node *node[K, V]) {
	// [] - [] - [x]
	if node == c.tail {
		return
	}

	defer func() { c.tail = node }()
	// [] - [x] - []
	if node.next != nil && node.previous != nil {
		node.previous.next = node.next
		node.next.previous = node.previous
		c.tail.next = node
		node.previous = c.tail
		node.next = nil
		return
	}
	// [x] - [] - []
	c.head = node.next
	node.next.previous = nil
	node.next = nil
	node.previous = c.tail
	c.tail.next = node

}

/*
NewCache init new cache. Key must be comparable

	capacity: this is the capacity of the lru cache. If the number of added elements is greater than the capacity the last element is removed

	nodeLifeTime: cache item lifetime
*/
func NewCache[K comparable, V any](capacity int, nodeLifeTime time.Duration) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		data:     make(map[K]*node[K, V], capacity),
		lifeTime: nodeLifeTime,
	}
}

// CacheSize number of items in the cache at the moment
func (c *Cache[K, V]) CacheSize() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.data)
}

// CacheCapacity current cache capacity
func (c *Cache[K, V]) CacheCapacity() int {
	return c.capacity
}

// Add add the element in cache. will return true if a new element was added.
// If an element was updated returns false.
func (c *Cache[K, V]) Add(key K, value V) bool {
	start := time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()

	// update exist element
	if cEl, ok := c.data[key]; ok {
		cEl.value = value
		cEl.time = start
		c.unsafeMoveToTail(cEl)
		return false

	}

	// delete of the latter when the size is exceeded
	if len(c.data) >= c.capacity {
		c.unsafeDelete(c.head)
	}

	c.unsafeAddToTail(&node[K, V]{key: key, value: value, time: start})
	return true
}

// Get get from cache by key. Return true if value in cache
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.data[key]

	if !ok {
		return c.zeroVal, false
	}

	if time.Since(node.time) > c.lifeTime {
		c.unsafeDelete(node)
		return c.zeroVal, false
	}

	c.unsafeMoveToTail(node)
	return node.value, true
}
