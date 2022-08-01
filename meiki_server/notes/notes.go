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
	ID    string `json:"id"`
	Title string `json:"title"`
	// TODO: Add useful timestamps
}

type NoteContentResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NotesStore struct {
	coll *mongo.Collection
}

var (
	ErrNoteDoesNotExist = errors.New("note does not exist")
	ErrInvalidId        = errors.New("invalid id")
	ErrInvalidTitle     = errors.New("invalid title")
)

func CreateNotesStore(ctx context.Context, coll *mongo.Collection) (NotesStore, error) {
	return NotesStore{coll}, nil
}

func validateNoteTitle(title string) error {
	if len(title) == 0 {
		return ErrInvalidTitle
	}

	return nil
}

func (ns NotesStore) Create(ctx context.Context, note Note) (string, error) {
	err := validateNoteTitle(note.Title)

	if err != nil {
		log.Error("invalid note", zap.Error(err))
		return "", err
	}

	result, err := ns.coll.InsertOne(ctx, note)

	if err != nil {
		log.Error("unable to create note", zap.Error(err))
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (ns NotesStore) List(ctx context.Context, username string) ([]NoteResponse, error) {
	cursor, err := ns.coll.Find(ctx, bson.M{"username": username})

	if err != nil {
		log.Error("unable to list notes", zap.Error(err))
		return nil, err
	}

	noteInfoList := []NoteResponse{}

	for cursor.Next(ctx) {
		var note Note

		err := cursor.Decode(&note)

		if err != nil {
			log.Error("unable to retrive note", zap.Error(err))
			return nil, err
		}

		noteInfoList = append(noteInfoList, NoteResponse{
			ID:    note.ID.Hex(),
			Title: note.Title,
		})
	}

	return noteInfoList, nil
}

func (ns NotesStore) Read(ctx context.Context, id string, username string) (NoteContentResponse, error) {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("invalid id requested on reading note", zap.Error(err))
		return NoteContentResponse{}, ErrInvalidId
	}

	var note Note
	err = ns.coll.FindOne(ctx, bson.M{"_id": docID, "username": username}).Decode(&note)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return NoteContentResponse{}, ErrNoteDoesNotExist
	}

	if err != nil {
		log.Error("unable to read note", zap.Error(err))
		return NoteContentResponse{}, err
	}

	return NoteContentResponse{
		Title:   note.Title,
		Content: note.Content,
	}, nil
}

func (ns NotesStore) Update(ctx context.Context, id string, username string, content string) error {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("invalid id requested on updating note", zap.Error(err))
		return ErrInvalidId
	}

	result, err := ns.coll.UpdateOne(ctx,
		bson.M{"_id": docID, "username": username},
		bson.M{"$set": bson.M{"content": content}},
	)

	if err != nil {
		log.Error("unable to update note")
		return err
	}

	if result.MatchedCount == 0 {
		log.Warn("note was not found for update request")
		return ErrNoteDoesNotExist
	}

	return nil
}

func (ns NotesStore) Rename(ctx context.Context, id string, username string, title string) error {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("invalid id requested on updating note", zap.Error(err))
		return ErrInvalidId
	}

	err = validateNoteTitle(title)

	if err != nil {
		log.Error("invalid note", zap.Error(err))
		return err
	}

	result, err := ns.coll.UpdateOne(ctx,
		bson.M{"_id": docID, "username": username},
		bson.M{"$set": bson.M{"title": title}},
	)

	if err != nil {
		log.Error("unable to update note title")
		return err
	}

	if result.MatchedCount == 0 {
		log.Warn("note was not found for rename request")
		return ErrNoteDoesNotExist
	}

	return nil
}

func (ns NotesStore) Delete(ctx context.Context, id string, username string) error {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Warn("invalid id requested on deleting note", zap.Error(err))
		return ErrInvalidId
	}

	result, err := ns.coll.DeleteOne(ctx, bson.M{"_id": docID, "username": username})

	if err != nil {
		log.Error("unable to delete note")
		return err
	}

	if result.DeletedCount == 0 {
		log.Warn("note was not found for delete request")
		return ErrNoteDoesNotExist
	}

	return nil
}

func (ns NotesStore) Search(ctx context.Context, query string, username string) ([]NoteResponse, error) {
	// TODO has common code with List

	cursor, err := ns.coll.Find(ctx,
		bson.M{
			"username": username,
			"$or": bson.A{
				bson.M{"title": bson.M{"$regex": query, "$options": "i"}},
				bson.M{"content": bson.M{"$regex": query, "$options": "i"}},
			},
		},
	)

	if err != nil {
		log.Error("unable to search notes", zap.Error(err))
		return nil, err
	}

	noteInfoList := []NoteResponse{}

	for cursor.Next(ctx) {
		var note Note

		err := cursor.Decode(&note)

		if err != nil {
			log.Error("unable to search note", zap.Error(err))
			return nil, err
		}

		noteInfoList = append(noteInfoList, NoteResponse{
			ID:    note.ID.Hex(),
			Title: note.Title,
		})
	}

	return noteInfoList, nil
}
