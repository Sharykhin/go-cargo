package main

import (
	"Sharykhin/go-cargo/infrastructure/http"
	"flag"
	"fmt"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", ":8000", "web server addr")
	flag.Parse()

	fmt.Printf("Run a web server on %s\n", addr)
	http.ListenAndServe(addr)
}
