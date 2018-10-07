package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var (
	bufferStorage bytes.Buffer
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	go bufferStorage.Write(body)
	//TODO: save to DB

	successResponse(w, nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}

func successResponse(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}
