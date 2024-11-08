package evictionPolicies

import (
	"enterpret/backend/shared"
)

type LFUEvictionPolicy struct {
	order           map[string]*shared.CacheItem
	keyTofrequency  map[string]int
	frequencyToKeys map[int]*shared.OrderedMap
	minimumFreq     int
}

func NewLFUEvictionPolicy() *LFUEvictionPolicy {
	// Initialize the order and frequencies maps
	return &LFUEvictionPolicy{
		order:           make(map[string]*shared.CacheItem),
		keyTofrequency:  make(map[string]int),
		frequencyToKeys: make(map[int]*shared.OrderedMap),
	}
}

func (p *LFUEvictionPolicy) OnAdd(item *shared.CacheItem) {
	// Add the item to the order map
	p.order[item.Key] = item
	p.keyTofrequency[item.Key] = 1
	if _, exists := p.frequencyToKeys[1]; !exists {
		p.frequencyToKeys[1] = shared.NewOrderedMap()
	}
	p.frequencyToKeys[1].Set(item.Key, true)
	p.minimumFreq = 1
}

func (p *LFUEvictionPolicy) OnAccess(item *shared.CacheItem) {
	// Increment the frequency of the accessed item
	if _, exists := p.keyTofrequency[item.Key]; exists {
		freq := p.keyTofrequency[item.Key]
		p.keyTofrequency[item.Key] = freq + 1
		p.frequencyToKeys[freq].Delete(item.Key)
		if _, exists := p.frequencyToKeys[freq+1]; !exists {
			p.frequencyToKeys[freq+1] = shared.NewOrderedMap()
		}
		p.frequencyToKeys[freq+1].Set(item.Key, true)
		if freq == p.minimumFreq && len(p.frequencyToKeys[freq].Values) == 0 {
			p.minimumFreq++
		}
	}
}

func (p *LFUEvictionPolicy) OnEvict() *shared.CacheItem {
	// Check if the order map is empty
	if len(p.order) == 0 {
		return nil
	}

	leastFreqKey := p.frequencyToKeys[p.minimumFreq].FirstElement()
	if leastFreqKey == "" {
		return nil
	}

	itemToEvict := p.order[leastFreqKey]

	delete(p.order, leastFreqKey)
	delete(p.keyTofrequency, leastFreqKey)
	p.frequencyToKeys[p.minimumFreq].Delete(leastFreqKey)

	if len(p.frequencyToKeys[p.minimumFreq].Values) == 0 {
		delete(p.frequencyToKeys, p.minimumFreq)
		p.minimumFreq++
	}

	return itemToEvict
}
