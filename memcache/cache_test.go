package memcache

import (
	"testing"
	"time"
)

func TestMemCache_Get(t *testing.T) {
	cache := NewMemCache(10 * time.Second)
	cacheWithValue := NewMemCache(10 * time.Second)
	cacheWithValue.Set("valid-key", "VALID data μ")
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		m       *MemCache
		args    args
		want    string
		wantErr bool
	}{
		{"Get empty key", cache, args{"emptykey"}, "", true},
		{"Get valid key", cacheWithValue, args{"valid-key"}, "VALID data μ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MemCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Set(t *testing.T) {
	cache := NewMemCache(10 * time.Second)

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		m       *MemCache
		args    args
		wantErr bool
	}{
		{"Set new key", cache, args{"newkey", "data"}, false},
		{"Set existing key", cache, args{"newkey", "data"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("MemCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemCache_handleTimeouts(t *testing.T) {
	cache := NewMemCache(2 * time.Second)
	key, value := "newkey", "newdata"

	type args struct {
		sleepTime time.Duration
	}
	tests := []struct {
		name    string
		m       *MemCache
		args    args
		toSet   bool
		want    string
		wantErr bool
	}{
		{"Set new key", cache, args{0 * time.Second}, true, "", false},
		{"Get key within ttl", cache, args{3 * time.Second}, false, value, false},
		{"Get key after ttl", cache, args{0 * time.Second}, false, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.toSet {
				if err := cache.Set(key, value); (err != nil) != tt.wantErr {
					t.Errorf("MemCache.Set() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {

				got, err := cache.Get(key)
				if (err != nil) != tt.wantErr {
					t.Errorf("MemCache.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("MemCache.Get() = %v, want %v", got, tt.want)
				}
			}
			time.Sleep(tt.args.sleepTime)
		})
	}
}
