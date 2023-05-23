package main

import (
	"flag"
	"laplace/core"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var staticDir string

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "Listen address")
	tls := flag.Bool("tls", true, "Use TLS")
	certFile := flag.String("certFile", staticDir + "/server.crt", "TLS cert file")
	keyFile := flag.String("keyFile", staticDir + "/server.key", "TLS key file")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	server := core.GetHttp(staticDir)

	if *tls {
		log.Println("Listening on TLS:", *addr)
		if err := http.ListenAndServeTLS(*addr, *certFile, *keyFile, server); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Listening:", *addr)
		if err := http.ListenAndServe(*addr, server); err != nil {
			log.Fatal(err)
		}
	}
}
