package notes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 5000000000*time.Millisecond)

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for notes test suite")
	}

	notes_db := client.Database("notesTest")
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

func newReqWithUserHeader(method string, route string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, route, body)
	req.Header.Set("X-Username", "alex")

	return req
}

func (s *NotesRoutesTestSuite) AssertListResponse(expectedTitleList []string) []notes.NoteResponse {
	req := newReqWithUserHeader("GET", "/list", nil)
	w := testhelpers.GetResponse(s.T(), s.router, req)

	assert.Equal(s.T(), 200, w.Code)

	var listResponse []notes.NoteResponse
	err := json.Unmarshal(w.Body.Bytes(), &listResponse)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(listResponse), len(expectedTitleList))

	for i := 0; i < len(listResponse); i += 1 {
		assert.Equal(s.T(), listResponse[i].Title, expectedTitleList[i])
	}

	return listResponse
}

func (s *NotesRoutesTestSuite) AssertAndGetListResponse(expectedTitleList []string) []notes.NoteResponse {
	// function exists for syntactic sugar
	return s.AssertListResponse(expectedTitleList)
}

func (s *NotesRoutesTestSuite) TestRoutesScenario() {
	// delete note invalid id
	req := newReqWithUserHeader("POST", "/delete/invalidID", nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Invalid id")

	// delete valid id but does not exist
	req = newReqWithUserHeader("POST", "/delete/"+primitive.NewObjectID().Hex(), nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// update valid id but does not exist
	updateRequest, _ := json.Marshal(notes.UpdateRequest{
		Content: "Modify the content",
	})

	req = newReqWithUserHeader("POST", "/update/"+primitive.NewObjectID().Hex(), bytes.NewBuffer(updateRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// read valid id but does not exist
	req = newReqWithUserHeader("GET", "/read/"+primitive.NewObjectID().Hex(), nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// create note and assert read
	createRequest, _ := json.Marshal(notes.CreateRequest{
		Title: "This is a new note",
	})

	req = newReqWithUserHeader("POST", "/create", bytes.NewBuffer(createRequest))
	w := testhelpers.GetResponse(s.T(), s.router, req)
	assert.Equal(s.T(), w.Code, 200)

	// assert create using list
	s.AssertListResponse([]string{"This is a new note"})

	// create note and assert read
	createRequest2, _ := json.Marshal(notes.CreateRequest{
		Title: "This is another note",
	})

	req = newReqWithUserHeader("POST", "/create", bytes.NewBuffer(createRequest2))
	w = testhelpers.GetResponse(s.T(), s.router, req)
	assert.Equal(s.T(), w.Code, 200)

	// assert create using list
	listResponse := s.AssertAndGetListResponse([]string{"This is a new note", "This is another note"})

	// update note note and assert read
	updateRequest, _ = json.Marshal(notes.UpdateRequest{
		Content: "A content has been added to this note",
	})

	req = newReqWithUserHeader("POST", "/update/"+listResponse[0].ID, bytes.NewBuffer(updateRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Updated note")

	// read the note and assert updated content
	req = newReqWithUserHeader("GET", "/read/"+listResponse[0].ID, nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "A content has been added to this note") // TODO: make this response notes response instead of string

	// delete note and assert read note does not exist
	req = newReqWithUserHeader("POST", "/delete/"+listResponse[0].ID, nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Deleted note")

	// assert create using list
	s.AssertListResponse([]string{"This is another note"})

	// assert read other note exists
	req, _ = http.NewRequest("GET", "/read/"+listResponse[0].ID, nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")
}

func (s *NotesRoutesTestSuite) TestRoutesParseError() {
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, notes.MSG_PARSE_ERROR)

	req, _ = http.NewRequest("POST", "/update/1234", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, notes.MSG_PARSE_ERROR)
}
