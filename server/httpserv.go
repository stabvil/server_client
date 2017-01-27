package main

import (
	"encoding/json"
	"fmt"
	g "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"net/http"
)

var session *g.Session

type jsondata struct {
	data string
}
type jsondataDB struct {
	Name string `gorethink:"name"`
	Data string `gorethink:"data"`
}

func status(w http.ResponseWriter, r *http.Request) {
	slcB, _ := json.Marshal([]string{"server", "running"})
	w.Write([]byte(slcB))
}

//initialisation connect
func initD() {
	var err error

	session, err = g.Connect(g.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "test",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

//out all data in terminal
func all(w http.ResponseWriter, r *http.Request) {
	initD()
	rows, err := g.Table("table").Run(session)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var xz []jsondataDB
	err = rows.All(&xz)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(xz)

	//rows.Close()

}

//save data in rethinkdb db=test, table=table
func saveData(w http.ResponseWriter, r *http.Request) {
	initD()

	var prob jsondata
	prob.data = mux.Vars(r)["data"]

	//json.NewEncoder(w).Encode(prob)

	var data = map[string]interface{}{
		"Name": mux.Vars(r)["name"],
		"Data": prob.data,
	}

	//json.NewEncoder(w).Encode(data)
	//w.Write([]byte(prob.data))
	_, err := g.Table("table").Insert(data).RunWrite(session)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//w.Write([]byte(data))

}

//get data from rethinkdb db=test, table=table
func getData(w http.ResponseWriter, r *http.Request) {
	initD()

	res, err := g.Table("table").Filter(g.Row.Field("Name").Eq(mux.Vars(r)["name"])).Run(session)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer res.Close()

	var response jsondataDB
	err = res.One(&response)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	answer := jsondata{
		data: response.Data,
	}
	fmt.Println(response)
	//json.NewEncoder(w).Encode(answer)
	w.Write([]byte(answer.data))

	//	w.Write([]byte(mux.Vars(r)["name"]))
}

//create table in db=test (when we've just opened rethinkdb in docker)
func createTable(w http.ResponseWriter, r *http.Request) {

	initD()

	_, err := g.DB("test").TableCreate("table").RunWrite(session)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

}

func httpserver() {

	r := mux.NewRouter()
	r.HandleFunc("/status", status)
	r.HandleFunc("/createtable", createTable)
	r.HandleFunc("/all", all).Methods("GET")
	r.HandleFunc("/{name}/{data}", saveData).Methods("GET")
	r.HandleFunc("/{name}", getData).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}
