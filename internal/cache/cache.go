package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entires  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entires:  make(map[string]cacheEntry),
		interval: interval,
	}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			cache.reapLoop()
		}
	}()

	return cache
}

func (c *Cache) Add(url string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entires[url] = cacheEntry{createdAt: time.Now(), val: data}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	if _, exist := c.entires[url]; exist {
		return c.entires[url].val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	for k, v := range c.entires {
		if time.Since(v.createdAt) > c.interval {
			delete(c.entires, k)
		}
	}

}
