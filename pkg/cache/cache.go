package cache

import (
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

type Cache struct {
	*cache.Cache
	Expiration      time.Duration
	CleanupInterval time.Duration
}

// NewCache returns Cache instance
func NewCache() (*Cache, error) {
	var c = &Cache{}
	var err error

	c.Expiration, err = time.ParseDuration(os.Getenv("CACHE_EXPIRATION"))
	if err != nil {
		return nil, err
	}

	c.CleanupInterval, err = time.ParseDuration(os.Getenv("CACHE_CLEANUP_INTERVAL"))
	if err != nil {
		return nil, err
	}

	return c, nil
}
