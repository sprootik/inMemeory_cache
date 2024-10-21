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
	capacity int
	head     *node
	tail     *node
	data     map[string]*node
	mu       sync.Mutex
	lifeTime time.Duration
}

/*
NewCache init new cache.

	capacity: this is the capacity of the lru cache. If the number of added elements is greater than the capacity the last element is removed

	nodeLifeTime: cache item lifetime
*/
func NewCache(capacity int, nodeLifeTime time.Duration) *Cache {
	return &Cache{
		capacity: capacity,
		data:     make(map[string]*node, capacity),
		lifeTime: nodeLifeTime,
	}
}

// CacheSize number of items in the cache at the moment
func (c *Cache) CacheSize() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.data)
}

// CacheCapacity current cache capacity
func (c *Cache) CacheCapacity() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.capacity
}

// SetCapacity set cache capacity
func (c *Cache) SetCapacity(capacity int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.capacity = capacity
}

// Add add the element in cache. will return true if a new element was added or an element
// was updated and the timeout has expired. If an item with the specified key is in the cache and the timeout
// has not expired, returns false.
func (c *Cache) Add(key string, value any) bool {
	element := &node{key: key, value: value, time: time.Now()}

	c.mu.Lock()

	// delete of the latter when the size is exceeded
	if len(c.data) >= c.capacity {
		c.unsafeRemove(c.tail)
	}

	// delete expired element
	if cEl, ok := c.data[key]; ok {
		if time.Since(cEl.time) <= c.lifeTime {
			c.mu.Unlock()
			return false
		} else {
			c.unsafeRemove(cEl)
		}
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
	c.mu.Unlock()
	return true
}

// Get get from cache by key
func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()

	element, ok := c.data[key]

	if !ok {
		c.mu.Unlock()
		return nil, false
	}

	if time.Since(element.time) > c.lifeTime {
		c.unsafeRemove(element)
		c.mu.Unlock()
		return nil, false
	}

	c.mu.Unlock()
	return element.value, true
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
}
