package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {

	slcB, _ := json.Marshal([]string{"server", "running"})
	w.Write([]byte(slcB))
}

func httpserver() {

	r := mux.NewRouter()
	r.HandleFunc("/status", status)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}
