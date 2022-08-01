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
	Title:    "This is a note title",
	Content:  "I programmed you to pass this test",
}

var note2 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "alex",
	Title:    "This is another note title",
	Content:  "Boom! Big reveal! I turned myself into a test case!",
}

var note3 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "shnoo",
	Title:    "This is shnoo's note different",
	Content:  "What is my purpose?; You are a dummy for this test; Oh my god",
}

var note4 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "alex",
	Title:    "Is another alex notes",
	Content:  "this just sounds like fixtures with extra steps",
}

var note5 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "alex",
	Title:    "Rick's note",
	Content:  "We are two humble notes in search for adventure",
}

var note6 = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "shnoo",
	Title:    "Morty's note",
	Content:  "We are two humble notes in search for adventure",
}

var invalidNote = notes.Note{
	ID:       primitive.NilObjectID,
	Username: "shnoo",
	Title:    "",
	Content:  "This is an invalid note",
}

func getTitleList(notesResponseList []notes.NoteResponse) []string {
	var titleList []string

	for _, note := range notesResponseList {
		titleList = append(titleList, note.Title)
	}

	return titleList
}

func (s *NotesStoreTestSuite) TestShouldCreateAndReadNote() {
	id, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)
	_, err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	storeNoteContentResp, err := s.notesStore.Read(s.ctx, id, note1.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), note1.Content, storeNoteContentResp.Content)
	assert.Equal(s.T(), note1.Title, storeNoteContentResp.Title)

	_, err = s.notesStore.Read(s.ctx, id, "differentUser")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestCreateShouldError() {
	_, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, invalidNote)
	assert.ErrorIs(s.T(), err, notes.ErrInvalidTitle)

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
	noteTitleList := getTitleList(noteInfoList)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []string{"This is a note title", "This is another note title"}, noteTitleList)
}

func (s *NotesStoreTestSuite) TestShouldUpdateNote() {
	id, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	err = s.notesStore.Update(s.ctx, id, note1.Username, "Content has been modified")
	assert.Nil(s.T(), err)

	contentResponse, err := s.notesStore.Read(s.ctx, id, note1.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), note1.Title, contentResponse.Title)
	assert.Equal(s.T(), "Content has been modified", contentResponse.Content)

	err = s.notesStore.Update(s.ctx, id, "differentUser", "Content has been modified")
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
	id, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	err = s.notesStore.Rename(s.ctx, id, note1.Username, "Note title has been modified")
	assert.Nil(s.T(), err)

	notesList, err := s.notesStore.List(s.ctx, note1.Username)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "Note title has been modified", notesList[0].Title)

	err = s.notesStore.Update(s.ctx, id, "differentUser", "Note title has been modified")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestRenameShouldError() {
	err := s.notesStore.Rename(s.ctx, "Invalid id", "testUser", "Testing")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	err = s.notesStore.Rename(s.ctx, primitive.NewObjectID().Hex(), "testUser", "")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidTitle)

	s.cancel()
	err = s.notesStore.Rename(s.ctx, primitive.NewObjectID().Hex(), "testUser", "Testing")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesStoreTestSuite) TestShouldDeleteNote() {
	id1, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	id2, err := s.notesStore.Create(s.ctx, note2)
	assert.Nil(s.T(), err)

	err = s.notesStore.Delete(s.ctx, id1, "differentUser")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)

	err = s.notesStore.Delete(s.ctx, id1, note1.Username)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Read(s.ctx, id1, note1.Username)
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)

	_, err = s.notesStore.Read(s.ctx, id2, note2.Username)
	assert.Nil(s.T(), err)
}

func (s *NotesStoreTestSuite) TestDeleteShouldError() {
	err := s.notesStore.Delete(s.ctx, "Invalid id", "testUser")
	assert.ErrorIs(s.T(), err, notes.ErrInvalidId)

	s.cancel()
	err = s.notesStore.Delete(s.ctx, primitive.NewObjectID().Hex(), "testUser")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *NotesRoutesTestSuite) TestShouldSearchNote() {
	_, err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note2)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note4)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note5)
	assert.Nil(s.T(), err)

	_, err = s.notesStore.Create(s.ctx, note6)
	assert.Nil(s.T(), err)

	results, err := s.notesStore.Search(s.ctx, "title", "alex")
	noteTitleList := getTitleList(results)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(noteTitleList), 2)
	assert.Equal(s.T(), []string{"This is a note title", "This is another note title"}, noteTitleList)

	results, err = s.notesStore.Search(s.ctx, "programmed", "alex")
	noteTitleList = getTitleList(results)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(noteTitleList), 1)
	assert.Equal(s.T(), []string{"This is a note title"}, noteTitleList)

	results, err = s.notesStore.Search(s.ctx, "this", "alex")
	noteTitleList = getTitleList(results)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(noteTitleList), 3)
	assert.Equal(s.T(), []string{"This is a note title", "This is another note title", "Is another alex notes"}, noteTitleList)
}
