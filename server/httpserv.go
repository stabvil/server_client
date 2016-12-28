package main

import (
	"net/http"
	"github.com/gorilla/mux"	
	"encoding/json"
)



func status(w http.ResponseWriter, r *http.Request) {

	
	slcB, _ := json.Marshal([]string{"server","ok"})
	w.Write([]byte(slcB))
}




func main() {
	
	r := mux.NewRouter()
	r.HandleFunc("/status", status)
	http.Handle("/", r)
	
	http.ListenAndServe(":8000", nil)
	
}