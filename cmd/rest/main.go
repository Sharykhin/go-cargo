package main

import (
	"Sharykhin/go-cargo/infrastructure/http"
	"flag"
	"fmt"
)

var (
	addr string
)

func main()  {
	flag.StringVar(&addr, "addr", "http://localhost:8080", "web server addr")
	flag.Parse()

	fmt.Printf("Going to rur a web server on %s\n", addr)
	http.ListenAndServe(addr)
}
