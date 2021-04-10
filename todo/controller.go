package todo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

var todos map[string]Todo = make(map[string]Todo)

func Register(router *mux.Router) {
	router.HandleFunc("", getList).Methods(http.MethodGet)
	router.HandleFunc("", post).Methods(http.MethodPost)
	router.HandleFunc("/{id}", get).Methods(http.MethodGet)
	router.HandleFunc("/{id}", put).Methods(http.MethodGet)
	router.HandleFunc("/{id}", del).Methods(http.MethodGet)
}

func getList(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(todos)
}

func post(rw http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var todo Todo
	json.Unmarshal(body, &todo)
	todo.Id = xid.New().String()
	todos[todo.Id] = todo
	json.NewEncoder(rw).Encode(todo)
}

func get(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if todo, exists := todos[id]; exists {
		json.NewEncoder(rw).Encode(todo)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}

func put(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body, _ := ioutil.ReadAll(r.Body)
	var todo Todo
	json.Unmarshal(body, &todo)
	todos[id] = todo
	json.NewEncoder(rw).Encode(todo)
}

func del(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	delete(todos, id)
}
