package memcache

import "errors"

// MemCache ...
type MemCache struct {
	hashmap map[string]string
}

// NewMemCache ...
func NewMemCache() *MemCache {
	hashmap := make(map[string]string)
	return &MemCache{hashmap}
}

// Get ...
func (m *MemCache) Get(key string) (string, error) {

	if value, found := m.hashmap[key]; found {
		return value, nil
	}
	return "", errors.New("not found")
}

// Set ...
func (m *MemCache) Set(key string, value string) error {
	if value == "" {
		return errors.New("missing value")
	}
	m.hashmap[key] = value
	return nil
}
