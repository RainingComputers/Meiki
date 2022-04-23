package notes

import (
	"context"
	"errors"

	"github.com/RainingComputers/Meiki/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Note struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
	// TODO: Add useful timestamps
}

type NoteResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// TODO: Add useful timestamps
}

type NotesStore struct {
	coll *mongo.Collection
}

var (
	ErrNoteDoesNotExist = errors.New("note does not exist")
	ErrInvalidId        = errors.New("invalid id")
)

func CreateNotesStore(ctx context.Context, coll *mongo.Collection) (NotesStore, error) {
	return NotesStore{coll}, nil
}

func (ns NotesStore) Create(ctx context.Context, note Note) (NoteResponse, error) {
	result, err := ns.coll.InsertOne(ctx, note)

	if err != nil {
		log.Error("Unable to create note", zap.Error(err))
		return NoteResponse{}, err
	}

	return NoteResponse{
		ID:       result.InsertedID.(primitive.ObjectID).Hex(),
		Title:    note.Title,
		Username: note.Username,
	}, nil
}

func (ns NotesStore) List(ctx context.Context, username string) ([]NoteResponse, error) {
	cursor, err := ns.coll.Find(ctx, bson.M{"username": username})

	if err != nil {
		log.Error("Unable to list notes", zap.Error(err))
		return nil, err
	}

	noteInfoList := []NoteResponse{}

	for cursor.Next(ctx) {
		var note Note

		err := cursor.Decode(&note)

		if err != nil {
			log.Error("Unable to retrive note", zap.Error(err))
			return nil, err
		}

		noteInfoList = append(noteInfoList, NoteResponse{
			ID:       note.ID.Hex(),
			Title:    note.Title,
			Username: note.Username,
		})
	}

	return noteInfoList, nil
}

func (ns NotesStore) Read(ctx context.Context, id string) (string, error) {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("Invalid id requested on reading note", zap.Error(err))
		return "", ErrInvalidId
	}

	var note Note
	err = ns.coll.FindOne(ctx, bson.M{"_id": docID}).Decode(&note)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", ErrNoteDoesNotExist
	}

	if err != nil {
		log.Error("Unable to read note", zap.Error(err))
		return "", err
	}

	return note.Content, nil
}

func (ns NotesStore) Update(ctx context.Context, id string, content string) error {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("Invalid id requested on updating note", zap.Error(err))
		return ErrInvalidId
	}

	result, err := ns.coll.UpdateOne(ctx,
		bson.M{"_id": docID},
		bson.M{"$set": bson.M{"content": content}},
	)

	if err != nil {
		log.Error("Unable to update note")
		return err
	}

	if result.MatchedCount == 0 {
		log.Warn("Note was not found for update request")
		return ErrNoteDoesNotExist
	}

	return nil
}

func (ns NotesStore) Delete(ctx context.Context, id string) error {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("Invalid id requested on deleting note", zap.Error(err))
		return ErrInvalidId
	}

	result, err := ns.coll.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		log.Error("Unable to delete note")
		return err
	}

	if result.DeletedCount == 0 {
		log.Warn("Note was not found for delete request")
		return ErrNoteDoesNotExist
	}

	return nil
}
