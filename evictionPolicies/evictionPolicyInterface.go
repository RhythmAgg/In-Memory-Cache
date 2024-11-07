package evictionPolicies

import (
	"enterpret/backend/shared"
)

// EvictionPolicy defines the interface for eviction policies
type EvictionPolicy interface {
	OnAdd(item *shared.CacheItem)
	OnAccess(item *shared.CacheItem)
	OnEvict() *shared.CacheItem
}
