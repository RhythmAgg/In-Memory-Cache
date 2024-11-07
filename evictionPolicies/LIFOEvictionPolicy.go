package evictionPolicies

import (
	"enterpret/backend/common"
)

type LIFOEvictionPolicy struct {
	// maintain the order of items in a slice
	order []*common.CacheItem
}

func NewLIFOEvictionPolicy() *LIFOEvictionPolicy {
	// Initialize the order slice
	return &LIFOEvictionPolicy{order: make([]*common.CacheItem, 0)}
}

func (p *LIFOEvictionPolicy) OnAdd(item *common.CacheItem) {
	// Add the item to the end of the order slice
	p.order = append(p.order, item)
}

func (p *LIFOEvictionPolicy) OnAccess(item *common.CacheItem) {
	// No special action needed on access for LIFO
}

func (p *LIFOEvictionPolicy) OnEvict() *common.CacheItem {
	// Check if the order slice is empty and evict the last item
	if len(p.order) == 0 {
		return nil
	}
	evicted := p.order[len(p.order)-1]
	p.order = p.order[:len(p.order)-1]
	return evicted
}
