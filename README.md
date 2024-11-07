# In-Memory Cache
Designed and implemented an in-memory caching library for general use in Go. The interface feature of the Go language allows for easy implementation of libraries for general usecase that are easily extensible.

## Features
- A general purpose in memory key-value cache
- Implements many different eviction policies like LRU, LFU, FIFO, LIFO and random eviction. Easily extensible to other custom eviction policies as well
- Thread safe using mutex guard. The mutex lock prevents the race conditions by ensuring the operations exclusive access to the cache
- Allows to set time to live ( ttl ) for each resource and runs a background go-routine at regular intervals to clear expired resources.
- Allows cache persitence. The cache can be enabled to be saved in a json file at regular intervals during runtime. The saved cache-state can also be loaded at initialisation.

## Implementation
- `cache package` - Package exposed to the clients
    - `cache/cache.go` - Implements the cache interface and cache-operations like Get, Set, RemoveItem. Also implements the background go routines like TTLCleanup
    - `cache/cachePersistence.go` - Implements the cache persistence logic like saveCache, loadCache
- `evictionPolicies package` - Internally used for extending the custom eviction policies
    - `evictionPolicies/evictionPolicyInterface.go` - Implements the eviction policy base interface. Every policy needs to implement this interface.
    - `evictionPolicies/[Random, FIFO, ... LRU]EvictionPolicy.go` - Implements the individual policies.
-  `shared package` - Contains the structs shared between the 2 packages like cacheItem.
- `InMemoryCache_test.go` - Contains all the feature tests for each specific feature of the cache.

## Run Tests
```
    go test
```
Note: Individual tests can also be run like
```
    go test -run ^<Test function name>
```