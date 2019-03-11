package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(5*time.Minute, 10*time.Minute)
)

func GetCache(prefix string, key string) (interface{}, bool) {
	return c.Get(prefix + "-" + key)
}

func SetCache(prefix string, key string, value interface{}) {
	c.Set(prefix+"-"+key, value, cache.DefaultExpiration)
}
