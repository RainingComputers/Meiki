package notes

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	Username string
	Title    string // TODO create compound unique index between username and title
	Content  string
}

type NotesStore struct {
	coll *mongo.Collection
}

var (
	ErrNotImplemented    = errors.New("not implemented") // TODO: remove this error
	ErrNoteDoesNotExist  = errors.New("note does not exist")
	ErrNoteAlreadyExists = errors.New("note already exists")
)

func CreateNotesStore(ctx context.Context, coll *mongo.Collection) (NotesStore, error) {
	return NotesStore{coll}, nil
}

func (ns NotesStore) Create(note Note) error {
	return ErrNotImplemented
}

func (ns NotesStore) List(username string) ([]string, error) {
	return nil, ErrNotImplemented
}

func (ns NotesStore) Read(username string, title string) (string, error) {
	return "", ErrNotImplemented
}

func (ns NotesStore) Update(username string, title string) error {
	return ErrNotImplemented
}

func (ns NotesStore) Delete() error {
	return ErrNotImplemented
}
