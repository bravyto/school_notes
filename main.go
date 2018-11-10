package main

import (
	"html/template"
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/notes", getNotesHandler).Methods("GET")
	r.HandleFunc("/notes/new", newNoteHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", getNoteHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", updateNoteHandler).Methods("POST")
	r.HandleFunc("/notes", createNoteHandler).Methods("POST")
	return r
}

func main() {
	connString := "dbname=school_note sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/index.html")
  t.Execute(w, "")
}
