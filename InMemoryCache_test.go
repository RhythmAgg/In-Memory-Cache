package main

import (
	"enterpret/backend/cache"
	"os"
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {
	c := cache.NewCache(2, "FIFO", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)

	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, exists := c.Get("b"); !exists || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
}

func TestCacheEvictionPolicy(t *testing.T) {
	c := cache.NewCache(2, "FIFO", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	if _, exists := c.Get("a"); exists {
		t.Errorf("Expected 'a' to be evicted")
	}
	if val, exists := c.Get("b"); !exists || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestFIFOEvictionPolicyWithCache(t *testing.T) {
	// Create a cache with FIFO eviction policy and capacity of 2
	c := cache.NewCache(2, "FIFO", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)

	c.Set("c", 3)

	if _, exists := c.Get("a"); exists {
		t.Errorf("Expected 'a' to be evicted")
	}
	if val, exists := c.Get("b"); !exists || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestLIFOEvictionPolicyWithCache(t *testing.T) {
	// Create a cache with LIFO eviction policy and capacity of 2
	c := cache.NewCache(2, "LIFO", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	if _, exists := c.Get("b"); exists {
		t.Errorf("Expected 'b' to be evicted")
	}
	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestLRUEvictionPolicyWithCache(t *testing.T) {
	// Create a cache with LRU eviction policy and capacity of 2
	c := cache.NewCache(2, "LRU", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)

	c.Get("a")

	c.Set("c", 3)

	if _, exists := c.Get("b"); exists {
		t.Errorf("Expected 'b' to be evicted")
	}
	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}

	c.Get("a")
	c.Set("d", 4)

	if _, exists := c.Get("c"); exists {
		t.Errorf("Expected 'c' to be evicted")
	}
	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, exists := c.Get("d"); !exists || val != 4 {
		t.Errorf("Expected 4, got %v", val)
	}
}

func TestLFUEvictionPolicyWithCache(t *testing.T) {
	// Create a cache with LFU eviction policy and a capacity of 2
	c := cache.NewCache(2, "LFU", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)

	c.Get("a")

	c.Set("c", 3)

	if _, exists := c.Get("b"); exists {
		t.Errorf("Expected 'b' to be evicted, but it is still present")
	}
	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 'a' to be present with value 1, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 'c' to be present with value 3, got %v", val)
	}
}

func TestRandomEvictionPolicyWithCache(t *testing.T) {
	// Create a cache with Random eviction policy and capacity of 3
	c := cache.NewCache(3, "Random", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4)

	evictedItems := []string{"a", "b", "c"}
	evicted := false
	for _, item := range evictedItems {
		if _, exists := c.Get(item); !exists {
			evicted = true
			break
		}
	}

	if !evicted {
		t.Errorf("Expected one of the items 'a', 'b', or 'c' to be evicted, but none were")
	}
}

func TestClearCache(t *testing.T) {
	// Create a cache with a capacity of 3 and any eviction policy (e.g., FIFO)
	c := cache.NewCache(3, "FIFO", 0, 0, false)

	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	c.Clear()

	if _, exists := c.Get("a"); exists {
		t.Errorf("Expected 'a' to be cleared from the cache")
	}
	if _, exists := c.Get("b"); exists {
		t.Errorf("Expected 'b' to be cleared from the cache")
	}
	if _, exists := c.Get("c"); exists {
		t.Errorf("Expected 'c' to be cleared from the cache")
	}
}

func TestTTLCacheCleanup(t *testing.T) {
	// Create a new cache with a capacity of 2 and TTL cleanup interval of 1 second
	c := cache.NewCache(3, "FIFO", 1, 0, false)

	c.Set("a", 1, 2)
	c.Set("b", 2, 5)

	time.Sleep(2 * time.Second)

	if _, exists := c.Get("a"); exists {
		t.Errorf("Expected item 'a' to be removed, but it still exists")
	}

	if _, exists := c.Get("b"); !exists {
		t.Errorf("Expected item 'b' to still be in the cache, but it was removed")
	}
}

func TestCachePersistence(t *testing.T) {
	// Create a new cache with a capacity of 2 and Cache persistence enabled
	os.Remove("cache.json")
	c := cache.NewCache(2, "FIFO", 0, 1, false)

	c.Set("a", 1)

	time.Sleep(2 * time.Second)

	_, err := os.Stat("cache.json")
	if os.IsNotExist(err) {
		t.Errorf("Error opening file: %v", err)
	}
	os.Remove("cache.json")
}

func TestLoadSavedCache(t *testing.T) {
	os.Remove("cache.json")
	c := cache.NewCache(2, "FIFO", 0, 1, false)

	c.Set("a", 1)

	time.Sleep(2 * time.Second)

	// Create a new cache with a capacity of 2 and load the cache state from the json file
	c2 := cache.NewCache(2, "FIFO", 0, 0, true)

	if _, exists := c2.Get("a"); !exists {
		t.Errorf("Expected 'a' to be present in the cache")
	}
	os.Remove("cache.json")
}
