package main

import (
	"fmt"
	"net/http"
	"html/template"
	"gopkg.in/mgo.v2-unstable"
	"github.com/BartoshMaxim/fakeTrella/db"
)
type trella interface {

}

var boardCollection []db.Board

var statusCollection []db.Status

var taskCollection []db.Task

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/client/dist/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	m := make(map[string]trella)
	m["board"] = boardCollection
	m["status"] = statusCollection
	m["task"] = taskCollection
	t.ExecuteTemplate(w, "index", m)
}

func initBoard(s *mgo.Session) {
	s.DB("fakeTrella").C("Board").Find(nil).All(&boardCollection)
}

func initStatus(s *mgo.Session) {
	s.DB("fakeTrella").C("Status").Find(nil).All(&statusCollection)
}

func initTask(s *mgo.Session) {
	s.DB("fakeTrella").C("Task").Find(nil).All(&taskCollection)
}


func main() {
	fmt.Println("Listening on port :3000")

	session, err := mgo.Dial("localhost:27017")
	if (err != nil) {
		panic(err)
	} else {
		initBoard(session)
		initStatus(session)
		initTask(session)
	}
	for _, doc := range boardCollection {
		println(doc.Id)
	}

	for _, doc := range statusCollection {
		println(doc.Name)
	}

	for _, doc := range taskCollection {
		println(doc.Message)
	}

	//http.Handle("/", http.FileServer(http.Dir("./client/dist")))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
