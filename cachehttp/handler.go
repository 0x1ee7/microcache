package cachehttp

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/0x1EE7/microcache/memcache"
)

// CacheHandler ...
type CacheHandler struct {
	cache *memcache.MemCache
}

//NewCacheHandler ...
func NewCacheHandler() CacheHandler {
	memcache := memcache.NewMemCache()
	handler := CacheHandler{memcache}
	return handler
}

func (h CacheHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]
	if key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}
	switch req.Method {
	case http.MethodGet:
		// GET
		value, err := h.cache.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%v", value)
	case http.MethodPost:
		// POST
		value, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)
			return
		}
		err = h.cache.Set(key, string(value))
		if err != nil {
			http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
