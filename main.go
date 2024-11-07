package main

import (
	"enterpret/backend/cache"
	"fmt"
)

func main() {
	fifoCache := cache.NewCache(3, "FIFO", 0)
	lruCache := cache.NewCache(3, "LRU", 0)

	fifoCache.Set("a", 1)
	fifoCache.Set("b", 2)
	fifoCache.Set("c", 3)
	fifoCache.Set("d", 4)

	fmt.Print(fifoCache.Get("b"))

	lruCache.Set("a", 1)
	lruCache.Set("b", 2)
	lruCache.Set("c", 3)
	lruCache.Get("a")
	lruCache.Set("d", 4)
}
