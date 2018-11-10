package main

import (
	"database/sql"
	"time"
)

type Store interface {
	CreateNote(note *Note) (int, error)
	GetNotes() ([]*Note, error)
	GetNote(id string) (*Note, error)
	UpdateNote(id string, note *Note) (error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateNote(note *Note) (int, error) {
	var id int
	err := store.db.QueryRow("INSERT INTO notes(title, text) VALUES ($1,$2) RETURNING id", note.Title, note.Text).Scan(&id)
	return id, err
}

func (store *dbStore) GetNotes() ([]*Note, error) {
	rows, err := store.db.Query("SELECT id, title, text, created_at, updated_at from notes ORDER BY updated_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []*Note{}
	for rows.Next() {
		note := &Note{}
		if err := rows.Scan(&note.Id, &note.Title, &note.Text, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (store *dbStore) GetNote(id string) (*Note, error) {
	row, err := store.db.Query("SELECT id, title, text, created_at, updated_at from notes where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	note := &Note{}
	for row.Next() {
		if err := row.Scan(&note.Id, &note.Title, &note.Text, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
	}
	return note, nil
}

func (store *dbStore) UpdateNote(id string, note *Note) (error) {
	_, err := store.db.Query("UPDATE notes SET title = $1, text = $2, updated_at = $3 WHERE id = $4", note.Title, note.Text, time.Now(), id)
	return err
}

var store Store

func InitStore(s Store) {
	store = s
}
