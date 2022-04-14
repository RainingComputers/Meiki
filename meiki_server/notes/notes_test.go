package notes_test

import (
	"context"
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
