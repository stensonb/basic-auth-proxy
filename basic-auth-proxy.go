package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/creack/goproxy"
	"github.com/goji/httpauth"
)

var port int
var configFilePath string

func init() {
	flag.IntVar(&port, "port", 9090, "port to bind to")
	flag.StringVar(&configFilePath, "config", "config.json", "path to the service config file")
}

// wrap http.ResponseWriter so we can extract status
type statusLoggingResponseWriter struct {
	status int
	http.ResponseWriter
}

func (w *statusLoggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
			clientIP = clientIP[:colon]
		}

		myW := &statusLoggingResponseWriter{http.StatusOK, w}

		startTime := time.Now()
		next.ServeHTTP(myW, r)
		finishTime := time.Now()

		user, _, ok := r.BasicAuth()
		if ok != true {
			user, _ = "", ""
		}

		log.Println(clientIP, r.Method, myW.status, r.RequestURI, user, finishTime.Sub(startTime))
	})
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
