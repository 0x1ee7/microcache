package main

import (
	"net/http"

	"github.com/0x1EE7/microcache/cachehttp"
)

func main() {
	handler := cachehttp.CacheHandler{}
	http.ListenAndServe(":8080", handler)
}
