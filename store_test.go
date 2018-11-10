package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=school_note_test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM notes")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateNote() {
	s.store.CreateNote(&Note{
		Title:  "test title",
		Text:   "test text",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM notes WHERE title='test title' AND text='test text'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetNote() {
	_, err := s.db.Query(`INSERT INTO notes (title, text) VALUES('title','text')`)
	if err != nil {
		s.T().Fatal(err)
	}

	notes, err := s.store.GetNotes()
	if err != nil {
		s.T().Fatal(err)
	}

	nNotes := len(notes)
	if nNotes != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nNotes)
	}

	expectedNote := Note{"title", "text"}
	if *notes[0] != expectedNote {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedNote, *notes[0])
	}
}
