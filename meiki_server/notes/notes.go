package notes

import (
	"context"
	"errors"

	"github.com/RainingComputers/Meiki/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Note struct {
	Username string
	Title    string
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
	mod := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}, {Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(ctx, mod)

	if err != nil {
		log.Error("Unable to create unique index in notes collection", zap.Error(err))
		return NotesStore{}, err
	}

	return NotesStore{coll}, nil
}

func (ns NotesStore) Create(ctx context.Context, note Note) error {
	_, err := ns.coll.InsertOne(ctx, note)

	if mongo.IsDuplicateKeyError(err) {
		log.Info("Duplicate note create attempt", zap.Error(err))
		return ErrNoteAlreadyExists
	}

	if err != nil {
		log.Error("Unable to create note", zap.Error(err))
		return err
	}

	return nil
}

func (ns NotesStore) List(ctx context.Context, username string) ([]string, error) {
	cursor, err := ns.coll.Find(ctx, bson.M{"username": username})

	if err != nil {
		log.Error("Unable to list notes", zap.Error(err))
		return nil, err
	}

	notesList := []string{}

	for cursor.Next(ctx) {
		var note Note

		err := cursor.Decode(&note)

		if err != nil {
			log.Error("Unable to retrive note", zap.Error(err))
			return nil, err
		}

		notesList = append(notesList, note.Title)
	}

	return notesList, nil
}

func (ns NotesStore) Read(ctx context.Context, username string, title string) (string, error) {
	var note Note
	err := ns.coll.FindOne(ctx, bson.M{"username": username}).Decode(&note)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", ErrNoteDoesNotExist
	}

	if err != nil {
		log.Error("Unable to read note", zap.Error(err))
		return "", err
	}

	return note.Content, nil
}

func (ns NotesStore) Update(ctx context.Context, username string, title string) error {
	return ErrNotImplemented
}

func (ns NotesStore) Delete(ctx context.Context) error {
	return ErrNotImplemented
}
