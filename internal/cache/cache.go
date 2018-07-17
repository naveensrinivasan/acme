package cache

import (
	c "github.com/patrickmn/go-cache"
)

// Cache interface provides caching options
type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, value interface{})
}
type cache struct {
	cache *c.Cache
}

// New returns the cache for storing certs
func New() Cache {
	return &cache{
		cache: c.New(-1, -1),
	}
}

// Get returns the cache item if it exists
func (c *cache) Get(k string) (interface{}, bool) {
	return c.cache.Get(k)
}

// Set - adds/updates the item to the cache
func (c *cache) Set(k string, value interface{}) {
	c.cache.Set(k, value, -1)
}
