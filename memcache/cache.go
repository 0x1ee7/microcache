package memcache

import (
	"errors"
	"time"
)

var ErrorNotFound = errors.New("not found")
var ErrorMisingValue = errors.New("missing value")
var ErrorNotModified = errors.New("not modified")

// MemCache ...
type MemCache struct {
	hashmap map[string]string
	channel chan string
}

// NewMemCache ...
func NewMemCache() *MemCache {
	hashmap := make(map[string]string)
	channel := make(chan string)
	cache := MemCache{hashmap, channel}
	go cache.handleTimeouts(channel)
	return &cache
}

// Get ...
func (m *MemCache) Get(key string) (string, error) {

	if value, found := m.hashmap[key]; found {
		return value, nil
	}
	return "", ErrorNotFound
}

// Set ...
func (m *MemCache) Set(key string, value string) error {
	if value == "" {
		return ErrorMisingValue
	}
	if _, exist := m.hashmap[key]; exist {
		return ErrorNotModified
	}

	go func() {
		<-time.After(5 * time.Second)
		m.channel <- key
	}()
	m.hashmap[key] = value
	return nil
}

func (m *MemCache) handleTimeouts(channel chan string) {
	for {
		key := <-channel
		delete(m.hashmap, key)
	}
}
