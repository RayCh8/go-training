package localcache

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	// mock time.Now()
	timeNow = func() time.Time {
		return time.Unix(1629446406, 0)
	}

	expect := &cache{
		es: map[string]*entry{
			"a": {
				val:    "123",
				expiry: timeNow().Add(30 * time.Second),
			},
		},
	}

	c := New()
	c.Set("a", "123")

	assert.Equal(t, reflect.DeepEqual(c, expect), true, "Should be equal")
}

func TestGet(t *testing.T) {
	c := &cache{
		es: map[string]*entry{
			"a": {
				val:    123,
				expiry: timeNow().Add(30 * time.Second),
			},
		},
	}

	val, err := c.Get("a")

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, val, 123, "Should be equal")
}

func TestGet_KeyNotExist(t *testing.T) {
	c := New()

	val, err := c.Get("a")

	assert.Nil(t, val)
	assert.Equal(t, err, ErrKeyNotExist, "Key doesn't exist in cache")
}

func TestGet_KeyExpiry(t *testing.T) {
	c := &cache{
		es: map[string]*entry{
			"a": {
				val:    123,
				expiry: timeNow().Add(30 * time.Second),
			},
		},
	}

	// mock time.Time.Before()
	timeBefore = func(_ time.Time, _ time.Time) bool {
		return true
	}
	val, err := c.Get("a")

	assert.Nil(t, val)
	assert.Equal(t, err, ErrKeyExpiry, "Key expiry")
}
