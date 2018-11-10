package main

import (
	"html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Note struct {
	Id				int					`json:"id"`
	Title  		string 			`json:"title"`
	Text   		string 			`json:"text"`
	CreatedAt *time.Time	`json:"created_at"`
	UpdatedAt *time.Time	`json:"updated_at"`
}

var notes []Note
var note Note
var savedNote Note

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := store.GetNotes()

	noteListBytes, err := json.Marshal(notes)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(noteListBytes)
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	note, err := store.GetNote(id)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t, _ := template.ParseFiles("assets/edit.html")
  t.Execute(w, note)
}

func newNoteHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/new.html")
  t.Execute(w, "")
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	note := Note{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	note.Title = r.Form.Get("title")
	note.Text = r.Form.Get("text")

	id, err := store.CreateNote(&note)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/notes/%d", id), http.StatusFound)
}

func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	note := Note{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	note.Title = r.Form.Get("title")
	note.Text = r.Form.Get("text")

	err = store.UpdateNote(id, &note)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/notes/%s", id), http.StatusFound)
}
