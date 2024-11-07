package evictionPolicies

import (
	"enterpret/backend/common"
	"math/rand"
)

type RandomEvictionPolicy struct {
	// Maintain the order of items in a slice
	order []*common.CacheItem
}

func NewRandomEvictionPolicy() *RandomEvictionPolicy {
	// Seed the random number generator to avoid predictable results
	return &RandomEvictionPolicy{order: make([]*common.CacheItem, 0)}
}

func (p *RandomEvictionPolicy) OnAdd(item *common.CacheItem) {
	// Add the item to the order slice
	p.order = append(p.order, item)
}

func (p *RandomEvictionPolicy) OnAccess(item *common.CacheItem) {
	// No action is required on access in a random eviction policy
}

func (p *RandomEvictionPolicy) OnEvict() *common.CacheItem {
	// Check if the order slice is empty
	if len(p.order) == 0 {
		return nil
	}

	randomIndex := rand.Intn(len(p.order))

	evicted := p.order[randomIndex]

	p.order[randomIndex] = p.order[len(p.order)-1]
	p.order = p.order[:len(p.order)-1]

	return evicted
}
