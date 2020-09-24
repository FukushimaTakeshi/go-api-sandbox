package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(r *http.Request) Response {
	return Respond(http.StatusOK, "Hello")
}

func TodoIndex(r *http.Request) Response {
	return Json(http.StatusOK, todos)
}

func TodoShow(r *http.Request) Response {
	id, _ := strconv.Atoi(mux.Vars(r)["todoId"])
	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		return Empty(http.StatusNotFound)
	}
	return Json(http.StatusOK, t)
}

func TodoCreate(r *http.Request) Response {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
	if err != nil {
		return Error(http.StatusInternalServerError, "request body is too large", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &todo); err != nil {
		return Error(http.StatusInternalServerError, "faild marshalling json", err)
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	return Created(http.StatusCreated, t, location)
}

func TodoDelete(r *http.Request) Response {
	id, _ := strconv.Atoi(mux.Vars(r)["todoId"])
	if err := RepoDestroyTodo(id); err != nil {
		return Empty(http.StatusNotFound)
	}
	return Empty(http.StatusNoContent)
}
