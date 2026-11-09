package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data             map[string]CacheEntry
	mu               sync.Mutex
	validIntervalSec time.Duration
	reapStarted      bool
}

type CacheEntry struct {
	data      []byte
	createdAt time.Time
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		data:      val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, has := c.data[key]

	if has {
		return val.data, has
	}
	return nil, has
}

func (c *Cache) ReapLoop() {
	c.mu.Lock()
	if c.reapStarted {
		c.mu.Unlock()
		return
	}
	c.mu.Unlock()

	ticker := time.NewTicker(c.validIntervalSec)

	for _ = range ticker.C {
		c.mu.Lock()
		deleteTime := time.Now().Add(-c.validIntervalSec * time.Second)
		for key, rec := range c.data {
			if rec.createdAt.After(deleteTime) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(validTimeSec time.Duration) *Cache {
	newCache := Cache{
		data:             map[string]CacheEntry{},
		mu:               sync.Mutex{},
		validIntervalSec: validTimeSec,
	}
	go newCache.ReapLoop()
	return &newCache
}
