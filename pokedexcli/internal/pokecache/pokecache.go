package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type PokeCache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(duration time.Duration) *PokeCache {
	pc := PokeCache{
		cache: make(map[string]cacheEntry),
	}

	go pc.reapLoop(duration)
	return &pc
}

func (pc *PokeCache) Add(key string, val []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (pc *PokeCache) Get(key string) ([]byte, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	entry, found := pc.cache[key]
	if !found {
		return nil, found
	}

	return entry.val, found
}

func (pc *PokeCache) reapLoop(duration time.Duration) {
	// ticker := time.NewTicker(time.Minute * 5)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	reap := func() {
		pc.mu.Lock()
		defer pc.mu.Unlock()

		now := time.Now()
		for k, v := range pc.cache {
			if age := now.Sub(v.createdAt); age > duration {
				delete(pc.cache, k)
			}
		}
	}

	for {
		<-ticker.C
		go reap()
	}
}
