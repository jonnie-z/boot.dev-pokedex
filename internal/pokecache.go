package internal

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	Interval     time.Duration
	Mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		CacheEntries: map[string]cacheEntry{},
		Interval:     interval,
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.CacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	entry, exists := c.CacheEntries[key]
	if exists {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			c.Mu.Lock()
			for key, cacheEntry := range c.CacheEntries {
				if cacheEntry.createdAt.Before(time.Now().Add(c.Interval * -1)) {
					delete(c.CacheEntries, key)
				}
			}
			c.Mu.Unlock()
		}
	}
}
