package evictionPolicies

import "enterpret/backend/shared"

type LRUEvictionPolicy struct {
	// Maintain the order of items in a slice
	order []*shared.CacheItem
}

func NewLRUEvictionPolicy() *LRUEvictionPolicy {
	// Initialize the order slice
	return &LRUEvictionPolicy{order: make([]*shared.CacheItem, 0)}
}

func (p *LRUEvictionPolicy) OnAdd(item *shared.CacheItem) {
	// Add the item to the end of the order slice. The first item in the slice is the most recently used
	p.order = append(p.order, item)
}

func (p *LRUEvictionPolicy) OnAccess(item *shared.CacheItem) {
	// Iterate over the order slice to find the item and move it to the front
	for i, it := range p.order {
		if it.Key == item.Key {
			p.order = append(p.order[:i], p.order[i+1:]...)
			p.order = append([]*shared.CacheItem{item}, p.order...)
			break
		}
	}
}

func (p *LRUEvictionPolicy) OnEvict() *shared.CacheItem {
	// Check if the order slice is empty and evict the last item (least recently used)
	if len(p.order) == 0 {
		return nil
	}
	evicted := p.order[len(p.order)-1]
	p.order = p.order[:len(p.order)-1]
	return evicted
}
