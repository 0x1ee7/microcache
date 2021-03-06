package cachehttp

import (
	"fmt"
	"net/http"
)

// CacheHandler ...
func CacheHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.Error(w, "Empty key not allowed", http.StatusBadRequest)
		return
	}
	switch req.Method {
	case http.MethodGet:

		fmt.Fprintf(w, "%v -> %v Get from cache\n", req.Method, req.URL.Path)
	case http.MethodPost:
		fmt.Fprintf(w, "%v -> %v Save to cache\n", req.Method, req.URL.Path)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
