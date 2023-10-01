package digicache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache() Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go cache.reapLoop(time.Minute * 30)

	return cache
}

func (c *Cache) Add(key string, data []byte) {
	c.mux.Lock()

	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		data:      data,
		createdAt: time.Now(),
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()

	defer c.mux.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()

	defer c.mux.Unlock()

	minutesAgo := time.Now().Add(-interval)
	for key, entry := range c.cache {
		if entry.createdAt.Before(minutesAgo) {
			delete(c.cache, key)
		}
	}
}
