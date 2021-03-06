package cachehttp

import (
	"fmt"
	"net/http"
)

// CacheHandler ...
type CacheHandler struct{}

func (h CacheHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]
	if key == "" {
		http.Error(w, "Empty key not allowed", http.StatusBadRequest)
		return
	}
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "%v -> %v Get from cache\n", req.Method, key)
	case http.MethodPost:
		fmt.Fprintf(w, "%v -> %v Save to cache\n", req.Method, key)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
