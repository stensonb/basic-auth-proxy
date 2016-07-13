package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

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
