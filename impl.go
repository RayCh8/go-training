package localcache

import (
	"errors"
	"sync"
	"time"
)

var (
	timeBefore = time.Time.Before
	timeNow    = time.Now
	// ErrKeyNotExist means that key doesn't exist.
	ErrKeyNotExist = errors.New("key does not exist.")
	// ErrKeyExpiry means key is expired.
	ErrKeyExpiry = errors.New("key expiry.")
)

type entry struct {
	val    interface{}
	expiry time.Time
}

type cache struct {
	l  sync.RWMutex
	es map[string]*entry
}

func (c *cache) Get(key string) (interface{}, error) {
	c.l.RLock()
	defer c.l.RUnlock()

	e, ok := c.es[key]

	if !ok {
		return nil, ErrKeyNotExist
	}

	if timeBefore(e.expiry, timeNow()) {
		delete(c.es, key)
		return nil, ErrKeyExpiry
	}

	return e.val, nil
}

func (c *cache) Set(key string, val interface{}) {
	c.l.Lock()
	defer c.l.Unlock()

	e := entry{
		val:    val,
		expiry: timeNow().Add(defaultTTL),
	}

	c.es[key] = &e
}

// New return Cache instance
func New() Cache {
	return &cache{
		es: make(map[string]*entry),
	}
}
