package main

import (
	"main/todo"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	handleRequest()
}

func handleRequest() {
	router := mux.NewRouter()
	router.Use(restApiHandler)
	todoRouter := router.PathPrefix("/todos").Subrouter()
	todo.Register(todoRouter)
	http.ListenAndServe(":8080", removeTrailingSlash(router))
}

// Remove trainling slashes (ex. /api/ -> /api) so that both paths can be handled in the same way.
// Regiters a middleware to remove trailing slashes not to mux.Router but to http module,
// because mux Router middlewares are fired only when matched routes are found.
// Note: mux.Router.StrictSlash(false) redirects (using 301 status code) /api/ to /api and vice versa,
// and mux doesn't provide a way to handle both paths without redirection.
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(rw, r)
	})
}

func restApiHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
