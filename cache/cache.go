package cache

import (
	"enterpret/backend/common"
	"enterpret/backend/evictionPolicies"
	"sync"
	"time"
)

// Cache is a thread-safe in-memory cache with support for custom eviction policies
type Cache struct {
	mu             sync.Mutex
	items          map[string]*common.CacheItem
	evictionPolicy evictionPolicies.EvictionPolicy
	capacity       int
}

func getEvictionPolicy(policy string) evictionPolicies.EvictionPolicy {
	switch policy {
	case "FIFO":
		return evictionPolicies.NewFIFOEvictionPolicy()
	case "LRU":
		return evictionPolicies.NewLRUEvictionPolicy()
	case "LIFO":
		return evictionPolicies.NewLIFOEvictionPolicy()
	case "LFU":
		return evictionPolicies.NewLFUEvictionPolicy()
	default:
		return evictionPolicies.NewFIFOEvictionPolicy()
	}
}

// StartTTLExpiryCleanup starts a background goroutine to clean up expired items periodically
func (c *Cache) StartTTLExpiryCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			c.mu.Lock()
			for key, item := range c.items {
				if time.Now().After(item.Expiration) {
					c.removeItem(key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

// NewCache creates a new Cache instance
func NewCache(capacity int, policy string, interval int) *Cache {
	cache := &Cache{
		items:          make(map[string]*common.CacheItem),
		evictionPolicy: getEvictionPolicy(policy),
		capacity:       capacity,
	}
	if interval > 0 {
		cache.StartTTLExpiryCleanup(time.Duration(interval) * time.Second)
	}
	return cache
}

// Set adds or updates an item in the cache
func (c *Cache) Set(key string, value interface{}, ttl ...int) {
	if len(ttl) == 0 {
		ttl = append(ttl, 1800)
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(time.Duration(ttl[0]) * time.Second)
	if item, exists := c.items[key]; exists {
		item.Value = value
		item.Timestamp = time.Now()
		item.Expiration = expiration
		c.evictionPolicy.OnAccess(item)
		return
	}

	item := &common.CacheItem{Key: key, Value: value, Timestamp: time.Now(), Expiration: expiration}

	// Check if the cache is full
	if len(c.items) >= c.capacity {
		evicted := c.evictionPolicy.OnEvict()
		if evicted != nil {
			c.removeItem(evicted.Key)
		}
	}
	c.items[key] = item
	c.evictionPolicy.OnAdd(item)
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.Expiration) {
		c.removeItem(key)
		return nil, false
	}

	c.evictionPolicy.OnAccess(item)
	return item.Value, true
}

// removeItem removes an item from the cache by key
func (c *Cache) removeItem(key string) {
	_, exists := c.items[key]
	if !exists {
		return
	}
	delete(c.items, key)
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*common.CacheItem)
}
