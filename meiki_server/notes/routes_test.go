package notes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)

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

func (s *NotesRoutesTestSuite) assertQueryResponse(expectedTitleList []string, queryType string) []notes.NoteResponse {
	req := newReqWithUserHeader("GET", queryType, nil)
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

func (s *NotesRoutesTestSuite) assertAndGetListResponse(expectedTitleList []string) []notes.NoteResponse {
	// function exists for syntactic sugar
	return s.assertQueryResponse(expectedTitleList, "/list")
}

func (s *NotesRoutesTestSuite) assertCreateResponse(w *httptest.ResponseRecorder) {
	assert.Equal(s.T(), w.Code, 200)

	var id string
	err := json.Unmarshal(w.Body.Bytes(), &id)
	assert.Nil(s.T(), err)

	assert.True(s.T(), len(id) > 5)
}

func (s *NotesRoutesTestSuite) assertReadResponse(w *httptest.ResponseRecorder, title string, content string) {
	assert.Equal(s.T(), w.Code, 200)

	var contentResponse notes.NoteContentResponse
	err := json.Unmarshal(w.Body.Bytes(), &contentResponse)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), title, contentResponse.Title)
	assert.Equal(s.T(), content, contentResponse.Content)
}

func (s *NotesRoutesTestSuite) TestRoutesScenario() {
	// delete note invalid id
	req := newReqWithUserHeader("DELETE", "/delete/invalidID", nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Invalid id")

	// delete valid id but does not exist
	req = newReqWithUserHeader("DELETE", "/delete/"+primitive.NewObjectID().Hex(), nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// update valid id but does not exist
	updateRequest, _ := json.Marshal(notes.UpdateRequest{
		Content: "Modify the content",
	})

	req = newReqWithUserHeader("PUT", "/update/"+primitive.NewObjectID().Hex(), bytes.NewBuffer(updateRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// read valid id but does not exist
	req = newReqWithUserHeader("GET", "/read/"+primitive.NewObjectID().Hex(), nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")

	// create note and assert read
	createRequest, _ := json.Marshal(notes.TitleRequest{
		Title: "This is a new note",
	})

	req = newReqWithUserHeader("POST", "/create", bytes.NewBuffer(createRequest))
	w := testhelpers.GetResponse(s.T(), s.router, req)
	s.assertCreateResponse(w)

	// assert create using list
	s.assertQueryResponse([]string{"This is a new note"}, "/list")

	// create note and assert read
	createRequest2, _ := json.Marshal(notes.TitleRequest{
		Title: "This is another note",
	})

	req = newReqWithUserHeader("POST", "/create", bytes.NewBuffer(createRequest2))
	w = testhelpers.GetResponse(s.T(), s.router, req)
	s.assertCreateResponse(w)

	s.assertQueryResponse([]string{"This is a new note", "This is another note"}, "/list")
	s.assertQueryResponse([]string{"This is another note"}, "/search?query=another")

	// assert create using list
	listResponse := s.assertAndGetListResponse([]string{"This is a new note", "This is another note"})

	// update note note and assert read
	updateRequest, _ = json.Marshal(notes.UpdateRequest{
		Content: "A content has been added to this note",
	})

	req = newReqWithUserHeader("PUT", "/update/"+listResponse[0].ID, bytes.NewBuffer(updateRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Updated note")

	// rename note and assert list response
	renameRequest, _ := json.Marshal(notes.TitleRequest{
		Title: "This is a new note with modified title",
	})

	req = newReqWithUserHeader("PUT", "/rename/"+listResponse[0].ID, bytes.NewBuffer(renameRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Renamed note")

	s.assertAndGetListResponse([]string{"This is a new note with modified title", "This is another note"})

	// read the note and assert updated content
	req = newReqWithUserHeader("GET", "/read/"+listResponse[0].ID, nil)
	w = testhelpers.GetResponse(s.T(), s.router, req)
	s.assertReadResponse(w, "This is a new note with modified title", "A content has been added to this note")

	// delete note and assert read note does not exist
	req = newReqWithUserHeader("DELETE", "/delete/"+listResponse[0].ID, nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Deleted note")

	// assert create using list
	s.assertQueryResponse([]string{"This is another note"}, "/list")

	// assert read other note does not exists
	req, _ = http.NewRequest("GET", "/read/"+listResponse[0].ID, nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Note does not exist")
}

func (s *NotesRoutesTestSuite) TestRoutesInputValidation() {
	createRequest, _ := json.Marshal(notes.TitleRequest{
		Title: "",
	})

	req := newReqWithUserHeader("POST", "/create", bytes.NewBuffer(createRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Invalid note title")

	renameRequest, _ := json.Marshal(notes.TitleRequest{
		Title: "",
	})

	req = newReqWithUserHeader("PUT", "/rename/"+primitive.NewObjectID().Hex(), bytes.NewBuffer(renameRequest))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "Invalid note title")
}

func (s *NotesRoutesTestSuite) TestRoutesParseError() {
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, notes.MSG_PARSE_ERROR)

	req, _ = http.NewRequest("PUT", "/rename/1234", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, notes.MSG_PARSE_ERROR)

	req, _ = http.NewRequest("PUT", "/update/1234", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, notes.MSG_PARSE_ERROR)
}
