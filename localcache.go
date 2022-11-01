package localcache

import (
	"errors"
	"time"
)

var (
	defaultTTL = time.Second * 30
	// ErrKeyNotExist means that key doesn't exist.
	ErrKeyNotExist = errors.New("key does not exist.")
	// ErrKeyExpiry means key is expiried.
	ErrKeyExpiry = errors.New("key expiry.")
)

// Cache is used to record data that user frequently fetch
type Cache interface {
	// Get is used to fetch value by key.
	Get(key string) (interface{}, error)
	// Set is used to set value by key
	Set(key string, val interface{})
}
