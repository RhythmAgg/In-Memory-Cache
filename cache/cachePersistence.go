package cache

import (
	"encoding/json"
	"enterpret/backend/shared"
	"os"
)

// SaveToDiskJSON saves the cache to disk using JSON encoding.
func SaveCacheToJSON(c *Cache, filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(c.items); err != nil {
		return err
	}
	return nil
}

// LoadFromDiskJSON loads the cache from disk using JSON encoding.
func LoadCacheFromJSON(c *Cache, filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	items := make(map[string]*shared.CacheItem)
	if err := decoder.Decode(&items); err != nil {
		return err
	}
	c.items = items
	return nil
}
