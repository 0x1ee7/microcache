package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0x1EE7/microcache/cachehttp"
	"github.com/0x1EE7/microcache/logging"
)

const defaultAddr = "localhost:8080"               // default webserver address
const defaultTTL = time.Duration(30 * time.Minute) // default cache ttl

var httpAddr = flag.String("http", defaultAddr, "HTTP service address")
var ttl = flag.Duration("ttl", defaultTTL, "Cache TTL 30m, 1h etc.")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: microcache -http=%v -ttl=%v\n", defaultAddr, ttl)
	fmt.Fprintf(os.Stderr, "env: CACHE_TTL overrides -ttl\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() > 0 {
		usage()
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	if eTTL := os.Getenv("CACHE_TTL"); eTTL != "" {
		envTTL, err := time.ParseDuration(eTTL)
		if err != nil {
			usage()
		}
		ttl = &envTTL
	}
	handler := cachehttp.NewCacheHandler(*ttl)
	logger.Printf("Server is starting... %v", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, logging.Handler(logger)(handler)))
}
