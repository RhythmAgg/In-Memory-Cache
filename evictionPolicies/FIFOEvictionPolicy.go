package evictionPolicies

import "enterpret/backend/common"

// FIFOEvictionPolicy is a FIFO eviction policy implementation using slices
type FIFOEvictionPolicy struct {
	order []*common.CacheItem
}

func NewFIFOEvictionPolicy() *FIFOEvictionPolicy {
	return &FIFOEvictionPolicy{order: make([]*common.CacheItem, 0)}
}

func (p *FIFOEvictionPolicy) OnAdd(item *common.CacheItem) {
	p.order = append(p.order, item)
}

func (p *FIFOEvictionPolicy) OnAccess(item *common.CacheItem) {
	// No special action needed on access for FIFO
}

func (p *FIFOEvictionPolicy) OnEvict() *common.CacheItem {
	if len(p.order) == 0 {
		return nil
	}
	evicted := p.order[0]
	p.order = p.order[1:]
	return evicted
}
