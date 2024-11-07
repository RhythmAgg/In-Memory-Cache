package main

import (
	"enterpret/backend/cache"
	"fmt"
)

func main() {
	fifoCache := cache.NewCache(3, "FIFO")
	lruCache := cache.NewCache(3, "LRU")

	fifoCache.Set("a", 1)
	fifoCache.Set("b", 2)
	fifoCache.Set("c", 3)
	fifoCache.Set("d", 4) // "a" will be evicted

	fmt.Print(fifoCache.Get("b"))

	lruCache.Set("a", 1)
	lruCache.Set("b", 2)
	lruCache.Set("c", 3)
	lruCache.Get("a")    // "a" becomes most recently used
	lruCache.Set("d", 4) // "b" will be evicted
}
