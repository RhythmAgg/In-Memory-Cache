package common

import "time"

type CacheItem struct {
	Key        string
	Value      interface{}
	Timestamp  time.Time
	Expiration time.Time
}
