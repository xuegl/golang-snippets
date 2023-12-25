package cache

import (
	"sync"
	"time"
)

// item represents a cache item with a value and an expiration time
type item[V any] struct {
	value  V
	expiry time.Time
}

// isExpired checks if the cache item has expired
func (i *item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

// TTLCache is a generic cache implementation with support for time-to-live(TTL) expiration
type TTLCache[K comparable, V any] struct {
	items   map[K]item[V]
	mu      sync.RWMutex
	ticker  *time.Ticker
	stopped bool
}

// NewTTLCache creates a new TTLCache instance and starts a goroutine to periodically remove
// expired items every 5 seconds
func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
	c := &TTLCache[K, V]{
		items:  make(map[K]item[V]),
		ticker: time.NewTicker(time.Second * 5),
	}

	go func() {
		for range c.ticker.C {
			c.mu.Lock()
			for k, i := range c.items {
				if i.isExpired() {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		}
	}()

	return c
}

// Set adds a new item to the cache with the specified key, value and ttl
func (c *TTLCache[K, V]) Set(k K, v V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stopped {
		panic("TTLCache has stopped")
	}
	c.items[k] = item[V]{
		value:  v,
		expiry: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with the given key from the cache
func (c *TTLCache[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.stopped {
		panic("TTLCache has stopped")
	}
	i, ok := c.items[k]
	if !ok {
		return i.value, false
	}

	if i.isExpired() {
		delete(c.items, k)
		return i.value, false
	}

	return i.value, true
}

// Remove removes the item with the specified key from the cache
func (c *TTLCache[K, V]) Remove(k K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stopped {
		panic("TTLCache has stopped")
	}
	delete(c.items, k)
}

// Pop removes and returns the item with the specified key from the cache
func (c *TTLCache[K, V]) Pop(k K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stopped {
		panic("TTLCache has stopped")
	}
	i, ok := c.items[k]
	if !ok {
		return i.value, false
	}

	delete(c.items, k)
	return i.value, !i.isExpired()
}

func (c *TTLCache[K, V]) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stopped {
		return
	}
	c.stopped = true
	c.ticker.Stop()
	c.items = make(map[K]item[V])
}
