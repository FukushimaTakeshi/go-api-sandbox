package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MyHandle func(*http.Request) Response

func IDShouldBeInt(h func(r *http.Request) Response, name string) MyHandle {
	return Logging(func(r *http.Request) Response {
		_, err := strconv.Atoi(mux.Vars(r)["todoId"])
		if err != nil {
			return Error(422, "todoId should be number", err)
		}
		return h(r)
	}, name)
}
