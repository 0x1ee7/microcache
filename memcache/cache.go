package memcache

import (
	"errors"
	"time"
)

// MemCache represents cache.
type MemCache struct {
	hashmap map[string]string
	channel chan string
}

// NewMemCache constructs a cache object.
// Starts a goroutine to handle timeout events.
func NewMemCache() *MemCache {
	hashmap := make(map[string]string)
	channel := make(chan string)
	cache := MemCache{hashmap, channel}
	go cache.handleTimeouts(channel)
	return &cache
}

// ErrorNotFound is returned when the key is not found the thehashmap
// ErrorMisingValue is returned when the value is empty
// ErrorNotModified is returned when the key is alre
var (
	ErrorNotFound    = errors.New("not found")
	ErrorMisingValue = errors.New("missing value")
	ErrorNotModified = errors.New("not modified")
)

// Get returns the value from the cache for a given key.
func (m *MemCache) Get(key string) (string, error) {

	if value, found := m.hashmap[key]; found {
		return value, nil
	}
	return "", ErrorNotFound
}

// Set saves the value to the cache for a given key. Also starts a goroutine to
// keep track of cache timeout if save is a success. Returns ErrorNotModified if
// cache already has a value for the key
func (m *MemCache) Set(key string, value string) error {
	if value == "" {
		return ErrorMisingValue
	}
	if _, found := m.hashmap[key]; found {
		return ErrorNotModified
	}

	m.hashmap[key] = value
	go func() {
		<-time.After(5 * time.Second)
		m.channel <- key
	}()
	return nil
}

func (m *MemCache) handleTimeouts(channel chan string) {
	for {
		key := <-channel
		delete(m.hashmap, key)
	}
}
