package evictionPolicies

import "enterpret/backend/shared"

// FIFOEvictionPolicy is a FIFO eviction policy implementation using slices
type FIFOEvictionPolicy struct {
	order []*shared.CacheItem
}

func NewFIFOEvictionPolicy() *FIFOEvictionPolicy {
	return &FIFOEvictionPolicy{order: make([]*shared.CacheItem, 0)}
}

func (p *FIFOEvictionPolicy) OnAdd(item *shared.CacheItem) {
	p.order = append(p.order, item)
}

func (p *FIFOEvictionPolicy) OnAccess(item *shared.CacheItem) {
	// No special action needed on access for FIFO
}

func (p *FIFOEvictionPolicy) OnEvict() *shared.CacheItem {
	if len(p.order) == 0 {
		return nil
	}
	evicted := p.order[0]
	p.order = p.order[1:]
	return evicted
}
