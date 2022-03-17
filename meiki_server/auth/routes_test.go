package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/RainingComputers/Meiki/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRoutesTestSuite struct {
	suite.Suite
	ctx    context.Context
	auth   auth.Auth
	cancel context.CancelFunc
	router *gin.Engine
}

func (s *AuthRoutesTestSuite) SetupTest() {
	log.Initialize()

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 1000000000*time.Millisecond)

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	auth_db := client.Database("auth")

	s.auth, err = auth.CreateAuth(s.ctx, auth_db.Collection("tokens"), auth_db.Collection("users"))
	assert.Nil(s.T(), err)

	s.router = gin.Default()
	auth.CreateRoutes(s.router, s.ctx, s.auth)
}

func TestAuthRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRoutesTestSuite))
}

func (s *AuthRoutesTestSuite) assertStatusCode(req *http.Request, expected int) {
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), expected, w.Code)
}

func (s *AuthRoutesTestSuite) TestRoutesScenario() {
	gin.SetMode(gin.ReleaseMode)
	s.auth.Delete(s.ctx, "alex")  // Just for cleanup
	s.auth.Delete(s.ctx, "shnoo") // Just for cleanup

	// TODO: Check response text as well

	// Create user
	credentialsBody, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-password",
	})

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(credentialsBody))
	s.assertStatusCode(req, 200)
	s.assertStatusCode(req, 400)

	// check auth status is false before logging in
	req, _ = http.NewRequest("POST", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", "randomToken")
	s.assertStatusCode(req, 401)

	// login and assert returned token headers and username
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(credentialsBody))
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), 200, w.Code)
	var sessionCredentials auth.SessionCredentials
	err := json.Unmarshal(w.Body.Bytes(), &sessionCredentials)
	assert.Nil(s.T(), err)

	token := sessionCredentials.Token
	assert.Equal(s.T(), sessionCredentials.Username, "alex")
	assert.True(s.T(), len(token) > 2)

	// check auth status is true after a login
	req, _ = http.NewRequest("POST", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	s.assertStatusCode(req, 200)

	// test login with wrong password
	badCredentialsBody1, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-wrong-password",
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody1))
	s.assertStatusCode(req, 401)

	// test login username does not exist
	badCredentialsBody2, _ := json.Marshal(auth.Credentials{
		Username: "shnoo",
		Password: "no-shnoo",
	})
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody2))
	s.assertStatusCode(req, 401)

	// test logout with bad token
	req, _ = http.NewRequest("POST", "/logout", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", "badToken")
	s.assertStatusCode(req, 400)

	// test logout with good token should log out
	req, _ = http.NewRequest("POST", "/logout", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	s.assertStatusCode(req, 200)

	// auth status should now return unauthorized
	req, _ = http.NewRequest("POST", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	s.assertStatusCode(req, 401)

	// delete user with bad creds should not delete anything
	req, _ = http.NewRequest("POST", "/delete", bytes.NewBuffer(badCredentialsBody1))
	s.assertStatusCode(req, 400) // TODO: Should be 401, along with another todo

	// delete user with good creds should delete user
	req, _ = http.NewRequest("POST", "/delete", bytes.NewBuffer(credentialsBody))
	s.assertStatusCode(req, 200)

	// delete user which doesn't exist should give 400
	req, _ = http.NewRequest("POST", "/delete", bytes.NewBuffer(credentialsBody))
	s.assertStatusCode(req, 400)
}
