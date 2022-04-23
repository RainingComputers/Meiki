package notes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/RainingComputers/Meiki/log"
	"github.com/RainingComputers/Meiki/notes"
	"github.com/RainingComputers/Meiki/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotesRoutesTestSuite struct {
	suite.Suite
	ctx        context.Context
	notesStore notes.NotesStore
	cancel     context.CancelFunc
	router     *gin.Engine
	coll       *mongo.Collection
}

func (s *NotesRoutesTestSuite) clean() {
	s.coll.DeleteMany(s.ctx, bson.M{})
}

func (s *NotesRoutesTestSuite) SetupTest() {
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

	s.router = gin.Default()
	authRouter := s.router.Group("/")
	notes.CreateRoutes(authRouter, s.ctx, s.notesStore)
}

func TestNotesRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(NotesRoutesTestSuite))
}

func (s *NotesRoutesTestSuite) TestRoutesScenario() {
	// delete note assert does not exist
	req, _ := http.NewRequest("POST", "/delete/invalidID", nil)
	testhelpers.AssertResponse(s.T(), s.router, req, 400, "Invalid id")

	// create note and assert read
	createRequest, _ := json.Marshal(notes.CreateRequest{
		Title: "This is a new note",
	})

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(createRequest))
	if err != nil {
		panic("shnoo says this should never happen!")
	}

	w := testhelpers.GetResponse(s.T(), s.router, req)
	assert.Equal(s.T(), w.Code, 200)

	//testhelpers.AssertResponse(s.T(), s.router, req, 200, "")

	// assert create using list
	req, _ = http.NewRequest("GET", "/list", nil)
	w = testhelpers.GetResponse(s.T(), s.router, req)

	assert.Equal(s.T(), 200, w.Code)
	var listResponse []notes.NoteResponse
	err = json.Unmarshal(w.Body.Bytes(), &listResponse)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(listResponse), 1)
	assert.Equal(s.T(), listResponse[0].Title, "This is a new note")

	// create note and assert read

	// list notes and assert

	// update note note and assert read

	//  delete note and assert read note does not exist

	// assert read other note exists
}
