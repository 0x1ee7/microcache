package memcache

import (
	"errors"
	"sync"
	"time"
)

// MemCache represents cache.
type MemCache struct {
	hashmap map[string]string
	ttl     time.Duration
	remove  chan string
	extend  chan string
	mu      sync.RWMutex
}

// NewMemCache constructs a cache object.
// Starts a goroutine to handle timeout events.
func NewMemCache(ttl time.Duration) *MemCache {
	hashmap := make(map[string]string)
	remove := make(chan string)
	extend := make(chan string)
	cache := MemCache{hashmap: hashmap, ttl: ttl, remove: remove, extend: extend}
	go cache.handleTimeouts(remove)
	return &cache
}

// ErrorNotFound is returned when the key is not found the thehashmap.
var ErrorNotFound = errors.New("not found")

// ErrorMisingValue is returned when the value is empty.
var ErrorMisingValue = errors.New("missing value")

// Get returns the value from the cache for a given key.
func (m *MemCache) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if value, found := m.hashmap[key]; found {
		return value, nil
	}
	return "", ErrorNotFound
}

// Set saves the value to the cache for a given key. Also starts a goroutine to
// keep track of cache timeout if save is a success.
func (m *MemCache) Set(key string, value string) error {
	if value == "" {
		return ErrorMisingValue
	}
	if _, err := m.Get(key); err == nil {
		m.extend <- key
	}
	m.mu.Lock()
	m.hashmap[key] = value
	m.mu.Unlock()
	go m.timeout(key)
	return nil
}

func (m *MemCache) handleTimeouts(remove chan string) {
	for {
		key := <-remove
		m.mu.Lock()
		delete(m.hashmap, key)
		m.mu.Unlock()
	}
}

func (m *MemCache) timeout(key string) {
	select {
	case <-m.extend:
		return
	case <-time.After(m.ttl):
		m.remove <- key
	}
}
