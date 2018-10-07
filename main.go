package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postHandler(w, r)
	case http.MethodGet:
		getHandler(w, r)
	default:
		defaultHandler(w, r)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
}

func getHandler(w http.ResponseWriter, r *http.Request) {

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}
