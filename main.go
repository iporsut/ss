package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "", "Host")
	flag.IntVar(&port, "port", 8080, "Port")
}

func main() {
	flag.Parse()

	http.Handle(
		"/",
		http.FileServer(http.Dir(".")),
	)
	log.Fatal(
		http.ListenAndServe(
			net.JoinHostPort(host, fmt.Sprint(port)),
			nil,
		),
	)
}
