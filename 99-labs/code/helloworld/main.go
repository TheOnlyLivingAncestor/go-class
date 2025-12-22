package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var hostname string
var version string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from %s running Go version %s!\n", hostname, version)
}

func main() {
	h, err := os.Hostname()
	version = runtime.Version()
	if err != nil {
		panic(err)
	}
	hostname = h
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
