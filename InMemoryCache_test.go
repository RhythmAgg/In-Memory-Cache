package main

import (
	"enterpret/backend/cache"
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	c := cache.NewCache(2, "FIFO")

	// Test setting values
	c.Set("a", 1)
	c.Set("b", 2)

	// Test retrieving values
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
	c.Set("c", 3) // "a" should be evicted in FIFO

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

	// Add items to the cache
	c.Set("a", 1)
	c.Set("b", 2)

	// Add another item to trigger eviction
	c.Set("c", 3)

	// Check if the first item "a" was evicted as per FIFO policy
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

	// Add items to the cache
	c.Set("a", 1)
	c.Set("b", 2)

	// Add another item to trigger eviction
	c.Set("c", 3)

	// Check if the last added item "b" was evicted as per LIFO policy
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

	// Add items to the cache
	c.Set("a", 1)
	c.Set("b", 2)

	// Access "a" to make it the most recently used
	c.Get("a")

	// Add another item to trigger eviction
	c.Set("c", 3)

	// Check if the least recently used item "b" was evicted as per LRU policy
	if _, exists := c.Get("b"); exists {
		t.Errorf("Expected 'b' to be evicted")
	}
	if val, exists := c.Get("a"); !exists || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, exists := c.Get("c"); !exists || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}

	// Verify the LRU order after another access
	// Access "a" again and add a new item to trigger eviction of "c"
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
