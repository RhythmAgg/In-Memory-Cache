package evictionPolicies

import "enterpret/backend/shared"

type LRUEvictionPolicy struct {
	// Maintain the order of items in a slice
	order *shared.OrderedMap
}

func NewLRUEvictionPolicy() *LRUEvictionPolicy {
	// Initialize the order slice
	return &LRUEvictionPolicy{order: shared.NewOrderedMap()}
}

func (p *LRUEvictionPolicy) OnAdd(item *shared.CacheItem) {
	// Add the item to the end of the order slice. The first item in the slice is the most recently used
	p.order.Set(item.Key, item)
}

func (p *LRUEvictionPolicy) OnAccess(item *shared.CacheItem) {
	// Iterate over the order slice to find the item and move it to the front
	if _, exists := p.order.Get(item.Key); exists {
		p.order.Delete(item.Key)
		p.order.Set(item.Key, item)
	}
}

func (p *LRUEvictionPolicy) OnEvict() *shared.CacheItem {
	// Check if the order slice is empty and evict the last item (least recently used)
	if len(p.order.Values) == 0 {
		return nil
	}
	key := p.order.FirstElement()
	item, _ := p.order.Get(key)
	p.order.Delete(key)
	return item.(*shared.CacheItem)
}
