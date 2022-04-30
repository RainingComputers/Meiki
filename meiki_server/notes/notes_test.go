package notes_test

import (
	"context"
	"testing"
	"time"

	"github.com/RainingComputers/Meiki/log"
	"github.com/RainingComputers/Meiki/notes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotesStoreTestSuite struct {
	suite.Suite
	ctx        context.Context
	notesStore notes.NotesStore
	cancel     context.CancelFunc
	coll       *mongo.Collection
}

func (s *NotesStoreTestSuite) clean() {
	s.coll.DeleteMany(s.ctx, bson.M{})
}

func (s *NotesStoreTestSuite) SetupTest() {
	log.Initialize()

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for notes test suite")
	}

	notes_db := client.Database("notes")
	s.coll = notes_db.Collection("notes")

	s.notesStore, err = notes.CreateNotesStore(s.ctx, s.coll)
	assert.Nil(s.T(), err)
	s.clean()
}

func (s *NotesStoreTestSuite) TearDownTest() {
	s.clean()
}

func TestNotesStoreTestSuite(t *testing.T) {
	suite.Run(t, new(NotesStoreTestSuite))
}

var note1 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "alex",
	Title:    "This is a note",
	Content:  "This is a test note, so it does not have many words in it",
}

var note2 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "alex",
	Title:    "This is another note",
	Content:  "You don't need to read this tho",
}

var note3 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "shnoo",
	Title:    "This is shnoo's note",
	Content:  "What is my purpose?; You are a dummy for this test; Oh my god",
}

func (s *NotesStoreTestSuite) TestShouldCreateAndReadNote() {
	noteInfo1, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)
	_, err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	storedContent, err := s.notesStore.Read(s.ctx, noteInfo1.ID, noteInfo1.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), note1.Content, storedContent)

	_, err = s.notesStore.Read(s.ctx, noteInfo1.ID, "differentUser")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestCreateShouldError() {
	_, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	s.cancel()
	_, err = s.notesStore.Create(s.ctx, note3)
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesStoreTestSuite) TestReadShouldError() {
	_, err := s.notesStore.Read(s.ctx, "Invalid id", "testUser")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	s.cancel()
	_, err = s.notesStore.Read(s.ctx, primitive.NewObjectID().Hex(), "testUser")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesStoreTestSuite) TestShouldListNotes() {

	notesList, err := s.notesStore.List(s.ctx, "alex")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []notes.NoteResponse{}, notesList)

	_, err = s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)
	_, err = s.notesStore.Create(s.ctx, note2)
	assert.Nil(s.T(), err)
	_, err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	noteInfoList, err := s.notesStore.List(s.ctx, "alex")
	noteTitleList := []string{noteInfoList[0].Title, noteInfoList[1].Title}
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []string{"This is a note", "This is another note"}, noteTitleList)
}

func (s *NotesStoreTestSuite) TestShouldUpdateNote() {
	noteInfo, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	err = s.notesStore.Update(s.ctx, noteInfo.ID, noteInfo.Username, "Content has been modified")
	assert.Nil(s.T(), err)

	content, err := s.notesStore.Read(s.ctx, noteInfo.ID, noteInfo.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "Content has been modified", content)

	err = s.notesStore.Update(s.ctx, noteInfo.ID, "differentUser", "Content has been modified")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestUpdateShouldError() {
	err := s.notesStore.Update(s.ctx, "Invalid id", "testUser", "Testing")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	s.cancel()
	err = s.notesStore.Update(s.ctx, primitive.NewObjectID().Hex(), "testUser", "Testing")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesStoreTestSuite) TestShouldRenameNote() {
	noteInfo, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	err = s.notesStore.Rename(s.ctx, noteInfo.ID, noteInfo.Username, "Note title has been modified")
	assert.Nil(s.T(), err)

	notesList, err := s.notesStore.List(s.ctx, noteInfo.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "Note title has been modified", notesList[0].Title)

	err = s.notesStore.Update(s.ctx, noteInfo.ID, "differentUser", "Note title has been modified")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestRenameShouldError() {
	err := s.notesStore.Rename(s.ctx, "Invalid id", "testUser", "Testing")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	s.cancel()
	err = s.notesStore.Rename(s.ctx, primitive.NewObjectID().Hex(), "testUser", "Testing")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesStoreTestSuite) TestShouldDeleteNote() {
	noteInfo1, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	noteInfo2, err := s.notesStore.Create(s.ctx, note2)
	assert.Nil(s.T(), err)

	err = s.notesStore.Delete(s.ctx, noteInfo1.ID, "differentUser")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)

	err = s.notesStore.Delete(s.ctx, noteInfo1.ID, noteInfo1.Username)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Read(s.ctx, noteInfo1.ID, noteInfo1.Username)
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)

	_, err = s.notesStore.Read(s.ctx, noteInfo2.ID, noteInfo2.Username)
	assert.Nil(s.T(), err)
}

func (s *NotesStoreTestSuite) TestDeleteShouldError() {
	err := s.notesStore.Delete(s.ctx, "Invalid id", "testUser")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	s.cancel()
	err = s.notesStore.Delete(s.ctx, primitive.NewObjectID().Hex(), "testUser")
	assert.ErrorIs(s.T(), err, context.Canceled)
}
