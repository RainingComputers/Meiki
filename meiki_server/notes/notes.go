package notes

import (
	"errors"
)

type Note struct {
	Username string
	Title    string // TODO create compound unique index between username and title
	Content  string
}

type NoteStore struct {
}

var (
	ErrNotImplemented    = errors.New("not implemented") // TODO: remove this error
	ErrNoteDoesNotExist  = errors.New("note does not exist")
	ErrNoteAlreadyExists = errors.New("note already exists")
)

func (ns NoteStore) Create() error {
	return ErrNotImplemented
}

func (ns NoteStore) List(username string) ([]string, error) {
	return nil, ErrNotImplemented
}

func (ns NoteStore) Read(username string, title string) (string, error) {
	return "", ErrNotImplemented
}

func (ns NoteStore) Update(username string, title string) error {
	return ErrNotImplemented
}

func (ns NoteStore) Delete() error {
	return ErrNotImplemented
}
