package main

import (
	grpc "Sharykhin/go-cargo/domain/company/port/grpc"
	"flag"
	"fmt"
	"log"
)

var (
	addr string
)

func main()  {
	flag.StringVar(&addr, "addr", ":50051", "grpc server addr")
	flag.Parse()

	fmt.Printf("Going to rur a web server on %s\n", addr)
	log.Fatal(grpc.ListenAndServe(addr))
}
