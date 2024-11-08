package shared

import "container/list"

type OrderedMap struct {
	Keys    *list.List               // Linked list for ordered Keys
	Values  map[string]interface{}   // Map for storing key-value pairs
	NodeMap map[string]*list.Element // Map for quick access to list elements
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		Keys:    list.New(),
		Values:  make(map[string]interface{}),
		NodeMap: make(map[string]*list.Element),
	}
}

func (o *OrderedMap) Set(key string, value interface{}) {
	if _, exists := o.Values[key]; !exists {
		// Add to linked list if it's a new key
		node := o.Keys.PushBack(key)
		o.NodeMap[key] = node
	}
	o.Values[key] = value
}

func (o *OrderedMap) Get(key string) (interface{}, bool) {
	value, exists := o.Values[key]
	return value, exists
}

func (o *OrderedMap) Delete(key string) {
	if _, exists := o.Values[key]; exists {
		// Delete the entry in Values map
		delete(o.Values, key)

		// Remove from linked list using node map
		if node, ok := o.NodeMap[key]; ok {
			o.Keys.Remove(node)
			delete(o.NodeMap, key)
		}
	}
}

func (o *OrderedMap) FirstElement() string {
	if o.Keys.Len() == 0 {
		return ""
	}
	key := o.Keys.Front().Value.(string)
	return key
}
