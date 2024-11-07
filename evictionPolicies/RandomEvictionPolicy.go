package evictionPolicies

import (
	"enterpret/backend/shared"
	"math/rand"
)

type RandomEvictionPolicy struct {
	// Maintain the order of items in a slice
	order []*shared.CacheItem
}

func NewRandomEvictionPolicy() *RandomEvictionPolicy {
	// Seed the random number generator to avoid predictable results
	return &RandomEvictionPolicy{order: make([]*shared.CacheItem, 0)}
}

func (p *RandomEvictionPolicy) OnAdd(item *shared.CacheItem) {
	// Add the item to the order slice
	p.order = append(p.order, item)
}

func (p *RandomEvictionPolicy) OnAccess(item *shared.CacheItem) {
	// No action is required on access in a random eviction policy
}

func (p *RandomEvictionPolicy) OnEvict() *shared.CacheItem {
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
