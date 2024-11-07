package evictionPolicies

import (
	"enterpret/backend/common"
)

// EvictionPolicy defines the interface for eviction policies
type EvictionPolicy interface {
	OnAdd(item *common.CacheItem)
	OnAccess(item *common.CacheItem)
	OnEvict() *common.CacheItem
}
