package main

import (
	"log"
	"net/http"
	"time"
)

var logger = func(method, uri, name string, status int, start time.Time) {
	log.Printf("\"method\":%q  \"uri\":%q    \"name\":%q   \"status\":%d \"time\":%q", method, uri, name, status, time.Since(start))
}

func Logging(h func(r *http.Request) Response, name string) MyHandle {
	return func(r *http.Request) Response {
		start := time.Now()
		result := h(r)
		logger(r.Method, r.URL.Path, name, result.Status(), start)
		return result
	}
}
