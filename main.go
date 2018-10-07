package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/korneliussamuel/go-sample-webserver/kit"

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
	http.HandleFunc("/{id:[0-9]+}", handler)
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
	url := r.URL.Path
	if hasDigit := regexp.MustCompile(`\d`).MatchString(url); hasDigit == false {
		responder.FailureResponse(w, http.StatusBadRequest)
		return
	}

	id := kit.GetIDFrom(url)

	person := getPersonFromBufferBy(id)
	if person != nil {
		responder.SuccessResponse(w, bufferStorage.Bytes())
		return
	}

	person = resource.FindPersonByID(DB, id)
	if person != nil {
		responder.SuccessResponse(w, kit.ToBytes(person))
		return
	}

	msg := fmt.Sprintf("data for user id %s does not exist", id)
	errMsg := resource.ErrMsg{msg}
	responder.SuccessResponse(w, kit.ToBytes(errMsg))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	responder.FailureResponse(w, http.StatusBadRequest)
}

func getPersonFromBufferBy(id string) *resource.Person {
	if bufferStorage.Len() != 0 {
		person := resource.Person{}
		json.Unmarshal(bufferStorage.Bytes(), person)

		if person.Id != id {
			return nil
		}

		return &person
	}

	return nil
}
