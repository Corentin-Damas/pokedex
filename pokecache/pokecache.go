package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	// !! 2 routines could hit the same memory alocation at the same time to prevent that we need a mutex
	mux *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()         // Lock access to other function in case an other fun try to access the cache when wetry to use it
	defer c.mux.Unlock() // Unlock the access when the funciton end
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
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
	xTimeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(xTimeAgo) {
			delete(c.cache, k)
		}
	}
}
