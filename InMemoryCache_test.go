package main

import (
	"enterpret/backend/cache"
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	c := cache.NewCache(2, "FIFO")

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
	c := cache.NewCache(2, "FIFO")

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
	c := cache.NewCache(2, "FIFO")

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
	c := cache.NewCache(2, "LIFO")

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
	c := cache.NewCache(2, "LRU")

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
	c := cache.NewCache(2, "LFU")

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

func TestClearCache(t *testing.T) {
	// Create a cache with a capacity of 3 and any eviction policy (e.g., FIFO)
	c := cache.NewCache(3, "FIFO")

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
