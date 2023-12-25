package cache

import "sync"

// Cache is a basic in-memory key-value cache implementation
type Cache[K comparable, V any] struct {
	items map[K]V
	mu    sync.RWMutex
}

// NewCache creates a new Cache instance
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]V),
	}
}

// Set adds or updates a key-value pair in the cache
func (c *Cache[K, V]) Set(k K, v V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[k] = v
}

// Get retrieves the value associated with the given key from the cache.
// The bool return value will be false if no matching key is found, and true otherwise
func (c *Cache[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[k]
	return v, ok
}

// Remove deletes the key-value pair with the specified key from the cache
func (c *Cache[K, V]) Remove(k K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, k)
}

// Pop removes and returns the value associated with the specified key from the cache
func (c *Cache[K, V]) Pop(k K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.items[k]
	if ok {
		delete(c.items, k)
	}
	return v, ok
}
