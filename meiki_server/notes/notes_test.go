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
	s.coll = notes_db.Collection("users")

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
	Username: "alex",
	Title:    "This is a note",
	Content:  "This is a test note, so it does not have many words in it",
}

var note2 = notes.Note{
	Username: "alex",
	Title:    "This is another note",
	Content:  "You don't need to read this tho",
}

var note3 = notes.Note{
	Username: "shnoo",
	Title:    "This is shnoo's note",
	Content:  "What is my purpose?; You are a dummy for this test; Oh my god",
}

func (s *NotesStoreTestSuite) TestShouldCreateAndReadNote() {
	err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)
	err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	storedContent, err := s.notesStore.Read(s.ctx, "alex", "This is a note")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), note1.Content, storedContent)
}

func (s *NotesStoreTestSuite) TestCreateShouldError() {
	err := s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)

	err = s.notesStore.Create(s.ctx, note1)
	assert.ErrorIs(s.T(), err, notes.ErrNoteAlreadyExists)
}

func (s *NotesStoreTestSuite) TestReadShouldError() {
	_, err := s.notesStore.Read(s.ctx, "someone", "Does not exist")
	assert.ErrorIs(s.T(), err, notes.ErrNoteDoesNotExist)
}

func (s *NotesStoreTestSuite) TestShouldListNotes() {

	notesList, err := s.notesStore.List(s.ctx, "alex")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []string{}, notesList)

	err = s.notesStore.Create(s.ctx, note1)
	assert.Nil(s.T(), err)
	err = s.notesStore.Create(s.ctx, note2)
	assert.Nil(s.T(), err)
	err = s.notesStore.Create(s.ctx, note3)
	assert.Nil(s.T(), err)

	notesList, err = s.notesStore.List(s.ctx, "alex")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []string{"This is a note", "This is another note"}, notesList)
}
