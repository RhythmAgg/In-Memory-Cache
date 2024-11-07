package evictionPolicies

import (
	"enterpret/backend/common"
)

type LFUEvictionPolicy struct {
	order       map[string]*common.CacheItem
	frequencies map[string]int
}

func NewLFUEvictionPolicy() *LFUEvictionPolicy {
	// Initialize the order and frequencies maps
	return &LFUEvictionPolicy{
		order:       make(map[string]*common.CacheItem),
		frequencies: make(map[string]int),
	}
}

func (p *LFUEvictionPolicy) OnAdd(item *common.CacheItem) {
	// Add the item to the order map
	p.order[item.Key] = item
	p.frequencies[item.Key] = 1
}

func (p *LFUEvictionPolicy) OnAccess(item *common.CacheItem) {
	// Increment the frequency of the accessed item
	if _, exists := p.frequencies[item.Key]; exists {
		p.frequencies[item.Key]++
	}
}

func (p *LFUEvictionPolicy) OnEvict() *common.CacheItem {
	// Check if the order map is empty
	if len(p.order) == 0 {
		return nil
	}

	// Find the item with the lowest frequency
	var leastFreqItem *common.CacheItem
	lowestFrequency := int(^uint(0) >> 1)

	for key, item := range p.order {
		if frequency, exists := p.frequencies[key]; exists && frequency < lowestFrequency {
			lowestFrequency = frequency
			leastFreqItem = item
		}
	}

	if leastFreqItem != nil {
		delete(p.order, leastFreqItem.Key)
		delete(p.frequencies, leastFreqItem.Key)
	}
	return leastFreqItem
}
