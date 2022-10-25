package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()

	got := len(c.es)

	assert.Equal(t, got, 0, "New cache size should be zero.")
}

func TestGet(t *testing.T) {
	c := New()

	c.Set("a", 123)
	got, ok := c.Get("a")
	size := len(c.es)

	assert.Nil(t, ok, "Should be nil")
	assert.Equal(t, size, 1, "Cache size should be one")
	assert.Equal(t, got, 123, "Should be equal")
}

func TestSet(t *testing.T) {
	c := New()

	c.Set("a", 123)
	c.Set("a", "456")
	got, ok := c.Get("a")

	assert.Equal(t, got, "456", "Should be equal")
	assert.Nil(t, ok, "Should be nil")
}

func TestKeyNotExist(t *testing.T) {
	c := New()

	got, ok := c.Get("a")

	assert.Nil(t, got)
	assert.Equal(t, ok, ErrKeyNotExist, "Key doesn't exist in cache")
}

func TestKeyExpiry(t *testing.T) {
	c := New()

	c.Set("a", 123)
	time.Sleep(30 * time.Second)
	got, ok := c.Get("a")

	assert.Nil(t, got)
	assert.Equal(t, ok, ErrKeyExpiry, "Key expiry")
}
