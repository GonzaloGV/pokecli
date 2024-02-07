package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	store map[string]cacheEntry
	mu    *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.store[key]

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(revalidationTime time.Duration) *Cache {
	cache := &Cache{store: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	go cache.reapLoop(revalidationTime)

	return cache
}
