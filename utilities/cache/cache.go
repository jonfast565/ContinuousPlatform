// a package for providing in memory caching
package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(5*time.Minute, 10*time.Minute)
)

// Get a cached object by its prefix and a key (for namespacing)
func GetCache(prefix string, key string) (interface{}, bool) {
	return c.Get(prefix + "-" + key)
}

// Set a cached object by its prefix, key, and value
func SetCache(prefix string, key string, value interface{}) {
	c.Set(prefix+"-"+key, value, cache.DefaultExpiration)
}
