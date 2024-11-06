package cache

import (
	dtos_cache "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/cache/dtos"
	"sync"
	"time"

	fscache "github.com/iqquee/fs-cache"
)

type ICacheUsersAdapter interface {
	Set(key string, value interface{}, duration time.Duration) error
	Get(key string) (interface{}, bool)
	Delete(key string)
	Clear()
}

type cacheUsersAdapter struct {
	cache *fscache.Operations
	mu    sync.Mutex
	items map[string]dtos_cache.CacheItem
}

func CacheUsersAdapter(cache *fscache.Operations) ICacheUsersAdapter {
	return &cacheUsersAdapter{
		cache: cache,
		items: make(map[string]dtos_cache.CacheItem),
	}
}

func (c *cacheUsersAdapter) Set(key string, value interface{}, duration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = dtos_cache.CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).Unix(),
	}

	return nil
}

func (c *cacheUsersAdapter) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found || time.Now().Unix() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

func (c *cacheUsersAdapter) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *cacheUsersAdapter) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]dtos_cache.CacheItem)
}
