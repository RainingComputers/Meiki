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

	"github.com/RainingComputers/Meiki/testhelpers"
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

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	auth_db := client.Database("auth")

	s.auth, err = auth.CreateAuth(s.ctx, auth_db.Collection("tokens"), auth_db.Collection("users"))
	assert.Nil(s.T(), err)

	s.router = gin.Default()
	authRouter := s.router.Group("/")
	auth.CreateRoutes(authRouter, s.ctx, s.auth)
}

func TestAuthRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRoutesTestSuite))
}

func (s *AuthRoutesTestSuite) assertSessionCredentials(req *http.Request, expectedUsername string) string {
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), 200, w.Code)

	var sessionCredentials auth.SessionCredentials
	err := json.Unmarshal(w.Body.Bytes(), &sessionCredentials)
	assert.Nil(s.T(), err)

	token := sessionCredentials.Token
	assert.Equal(s.T(), sessionCredentials.Username, expectedUsername)
	assert.True(s.T(), len(token) > 2)

	return token
}

func (s *AuthRoutesTestSuite) TestRoutesScenario() {
	gin.SetMode(gin.ReleaseMode)
	s.auth.Delete(s.ctx, "alex") // Cleanup before test
	s.auth.Delete(s.ctx, "shnoo")

	// create user
	credentialsBody, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-password",
	})

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(credentialsBody))
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "User successfully created")

	// test create existing user
	req, _ = http.NewRequest("POST", "/create", bytes.NewBuffer(credentialsBody))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "User already exists")

	// test login with wrong password
	badCredentialsBody1, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-wrong-password",
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody1))
	testhelpers.AssertResponseString(s.T(), s.router, req, 401, "Password does not match")

	// test login username does not exist
	badCredentialsBody2, _ := json.Marshal(auth.Credentials{
		Username: "shnoo",
		Password: "no-shnoo",
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody2))
	testhelpers.AssertResponseString(s.T(), s.router, req, 401, "User does not exist")

	// check auth status is false before logging in
	req, _ = http.NewRequest("GET", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", "randomToken")
	testhelpers.AssertResponseString(s.T(), s.router, req, 401, "User token does not exist")

	// login and assert returned token headers and username
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(credentialsBody))
	token := s.assertSessionCredentials(req, "alex")

	// check auth status is true after a login
	req, _ = http.NewRequest("GET", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "alex")

	// test logout with bad token
	req, _ = http.NewRequest("POST", "/logout", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", "badToken")
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "User token does not exist")

	// test logout with good token should log out
	req, _ = http.NewRequest("POST", "/logout", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "User logged out successfully")

	// auth status should now return unauthorized
	req, _ = http.NewRequest("GET", "/authStatus", nil)
	req.Header.Set("X-Username", "alex")
	req.Header.Set("X-Token", token)
	testhelpers.AssertResponseString(s.T(), s.router, req, 401, "Invalid or wrong credentials")

	// delete user with bad creds should not delete anything
	req, _ = http.NewRequest("DELETE", "/delete", bytes.NewBuffer(badCredentialsBody1))
	testhelpers.AssertResponseString(s.T(), s.router, req, 401, "Password does not match")

	// delete user with good creds should delete user
	req, _ = http.NewRequest("DELETE", "/delete", bytes.NewBuffer(credentialsBody))
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "User deleted user successfully")

	// delete user which doesn't exist should give 400
	req, _ = http.NewRequest("DELETE", "/delete", bytes.NewBuffer(credentialsBody))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, "User does not exist")
}

func (s *AuthRoutesTestSuite) TestRoutesInputValidation() {
	badUsernameCredentials, _ := json.Marshal(auth.Credentials{
		Username: "*a",
		Password: "something",
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(badUsernameCredentials))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_INVALID_USERNAME)

	req, _ = http.NewRequest("POST", "/create", bytes.NewBuffer(badUsernameCredentials))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_INVALID_USERNAME)

	badPasswordCredentials, _ := json.Marshal(auth.Credentials{
		Username: "shnoo",
		Password: "1234",
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badPasswordCredentials))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_INVALID_PASSWORD)

	req, _ = http.NewRequest("POST", "/create", bytes.NewBuffer(badPasswordCredentials))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_INVALID_PASSWORD)
}

func (s *AuthRoutesTestSuite) TestRoutesParseError() {
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_UNABLE_TO_PARSE_CREDENTIALS)

	req, _ = http.NewRequest("DELETE", "/delete", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_UNABLE_TO_PARSE_CREDENTIALS)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("test")))
	testhelpers.AssertResponseString(s.T(), s.router, req, 400, auth.MSG_UNABLE_TO_PARSE_CREDENTIALS)
}
