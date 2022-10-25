package localcache

import (
	"errors"
	"time"
)

var (
	// Time to live of cache key.
	DefaultTTL = time.Second * 30
	// Error message means that key does not exist.
	ErrKeyNotExist = errors.New("key does not exist.")
	// Error message means key is expiried.
	ErrKeyExpiry = errors.New("key expiry.")
)

// Cache is used to record data that user frequently fetch
type Cache interface {
	// Get is used to fetch value by key.
	Get(key string) (interface{}, error)
	// Set is used to set value by key
	Set(key string, val interface{})
}
