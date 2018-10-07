package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/korneliussamuel/go-sample-webserver/db"
	"github.com/korneliussamuel/go-sample-webserver/resource"
	"github.com/korneliussamuel/go-sample-webserver/responder"
)

var (
	bufferStorage bytes.Buffer
	DB            *sql.DB
)

func main() {
	DB = db.Setup()

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

	person := resource.Person{}
	if err := json.Unmarshal(body, &person); err != nil {
		panic(err)
	}
	go person.Save(DB)

	responder.SuccessResponse(w, nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}
