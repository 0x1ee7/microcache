package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0x1EE7/microcache/cachehttp"
	"github.com/0x1EE7/microcache/logging"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	handler := cachehttp.NewCacheHandler()
	logger.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8080", logging.Handler(logger)(handler)))
}
