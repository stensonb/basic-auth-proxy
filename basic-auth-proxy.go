package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/creack/goproxy"
	"github.com/goji/httpauth"
)

var port int
var configFilePath string

func init() {
	flag.IntVar(&port, "port", 9090, "port to bind to")
	flag.StringVar(&configFilePath, "config", "config.json", "path to the service config file")
}

func main() {
	flag.Parse()
	config := NewConfig(configFilePath)
	log.Println("Read config file from from", configFilePath)

	serviceHandler := http.HandlerFunc(goproxy.NewMultipleHostReverseProxy(config.Registry))
	authHandler := httpauth.SimpleBasicAuth(config.User, config.Password)

	http.Handle("/", logHandler(authHandler(serviceHandler)))

	log.Println(fmt.Sprintf("Listening on %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
