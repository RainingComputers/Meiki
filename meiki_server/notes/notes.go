package notes

import (
	"context"
	"errors"

	"github.com/RainingComputers/Meiki/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Note struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

type NoteInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
}

type NotesStore struct {
	coll *mongo.Collection
}

var (
	ErrNotImplemented    = errors.New("not implemented") // TODO: remove this error
	ErrNoteDoesNotExist  = errors.New("note does not exist")
	ErrNoteAlreadyExists = errors.New("note already exists")
	ErrInvalidId         = errors.New("invalid id")
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

func (ns NotesStore) Create(ctx context.Context, note Note) (NoteInfo, error) {
	result, err := ns.coll.InsertOne(ctx, note)

	if mongo.IsDuplicateKeyError(err) {
		log.Info("Duplicate note create attempt", zap.Error(err))
		return NoteInfo{}, ErrNoteAlreadyExists
	}

	if err != nil {
		log.Error("Unable to create note", zap.Error(err))
		return NoteInfo{}, err
	}

	return NoteInfo{
		ID:       result.InsertedID.(primitive.ObjectID).Hex(),
		Title:    note.Title,
		Username: note.Username,
	}, nil
}

func (ns NotesStore) List(ctx context.Context, username string) ([]NoteInfo, error) {
	cursor, err := ns.coll.Find(ctx, bson.M{"username": username})

	if err != nil {
		log.Error("Unable to list notes", zap.Error(err))
		return nil, err
	}

	noteInfoList := []NoteInfo{}

	for cursor.Next(ctx) {
		var note Note

		err := cursor.Decode(&note)

		if err != nil {
			log.Error("Unable to retrive note", zap.Error(err))
			return nil, err
		}

		noteInfoList = append(noteInfoList, NoteInfo{
			ID:       note.ID.String(),
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

	if result.MatchedCount == 0 {
		log.Warn("Note was not found for update request")
		return ErrNoteDoesNotExist
	}

	if err != nil {
		log.Error("Unable to update note")
		return err
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

	if result.DeletedCount == 0 {
		log.Warn("Note was not found for delete request")
		return ErrNoteDoesNotExist
	}

	if err != nil {
		log.Error("Unable to delete note")
		return err
	}

	return nil
}
