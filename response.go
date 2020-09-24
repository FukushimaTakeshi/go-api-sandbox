package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response interface {
	Write(w http.ResponseWriter)
	Status() int
}

type NormalResponse struct {
	status int
	body   []byte
	header http.Header
}

func (r *NormalResponse) Write(w http.ResponseWriter) {
	header := w.Header()
	for k, v := range r.header {
		header[k] = v
	}
	w.WriteHeader(r.status)
	w.Write(r.body)
}

func (r *NormalResponse) Status() int {
	return r.status
}

func (r *NormalResponse) Header(key, value string) *NormalResponse {
	r.header.Set(key, value)
	return r
}

func Empty(status int) *NormalResponse {
	return Respond(status, nil)
}

func Json(status int, body interface{}) *NormalResponse {
	return Respond(status, body).Header("Content-Type", "application/json")
}

func Created(status int, body interface{}, location string) *NormalResponse {
	return Json(status, body).Header("Location", location)
}

func Error(status int, message string, err error) *NormalResponse {
	log.Printf("%s, %s", message, err)
	return Respond(status, message).Header("Content-Type", "application/json")
}

func Respond(status int, body interface{}) *NormalResponse {
	var b []byte
	var err error
	switch t := body.(type) {
	case string:
		b = []byte(t)
	default:
		if b, err = json.Marshal(body); err != nil {
			return Error(http.StatusInternalServerError, "failed marshalling json", err)
		}
	}

	return &NormalResponse{
		status: status,
		body:   b,
		header: make(http.Header),
	}
}
