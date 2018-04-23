// Simple groupcache server running on :8080
package main

import (
	"log"
	"net/http"

	"github.com/brendanjryan/groupcache-bazel"
)

const (
	addr = ":8080"
)

func main() {
	pool := groupcache.NewHTTPPool(addr)

	log.Println("server listening on", addr)

	pool.Set(addr)
	http.ListenAndServe(addr, nil)
}
