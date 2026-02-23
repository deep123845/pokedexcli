package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:      &sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}

	go reapLoop(&cache, interval)

	return cache
}

func reapLoop(cache *Cache, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		cache.mu.Lock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) > interval {
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}
