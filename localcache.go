package localcache

import (
	"errors"
	"time"
)

var (
	// Time to live of cache key.
	DefaultTTL = time.Second * 30
	// Error message that key is not exist.
	ErrKeyNotExist = errors.New("key does not exist.")
	// Error message that key is expiry.
	ErrKeyExpiry = errors.New("key expiry.")
)

// Cache is used to record data that user frequently fetch
type Cache interface {
	// Get is used to fetch value by key.
	Get(key string) (interface{}, error)
	// Set is used to set value by key
	Set(key string, val interface{})
}
