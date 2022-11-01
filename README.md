# go-training

## localcache
```
type Cache interface {
	// Get is used to fetch value by key.
	Get(key string) (interface{}, error)
	// Set is used to set value by key
	Set(key string, val interface{})
}
```