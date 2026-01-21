package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	lock  *sync.Mutex
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		lock:  &sync.Mutex{},
		entry: map[string]cacheEntry{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	entry, ok := cache.entry[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		now := time.Now()
		cache.lock.Lock()

		for key, entry := range cache.entry {
			if d := now.Sub(entry.createdAt); d > interval { // dubious logic? extension says ok...
				delete(cache.entry, key)
			}
		}

		cache.lock.Unlock()
	}
}
