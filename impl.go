package localcache

import (
	"time"
)

type entry struct {
	val    interface{}
	expiry time.Time
}

type cache struct {
	es map[string]*entry
}

func (c *cache) Get(key string) (interface{}, error) {
	e, ok := c.es[key]

	if !ok {
		return nil, ErrKeyNotExist
	}

	if e.expiry.Before(time.Now()) {
		delete(c.es, key)
		return nil, ErrKeyExpiry
	}

	return e.val, nil
}

func (c *cache) Set(key string, val interface{}) {
	e := entry{
		val:    val,
		expiry: time.Now().Add(DefaultTTL),
	}

	c.es[key] = &e
}

// New return Cache instance
func New() cache {
	c := cache{
		es: make(map[string]*entry),
	}

	return c
}
