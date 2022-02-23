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

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)

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
	s.auth.Delete(s.ctx, "alex") // Just for cleanup

	credentialsBody, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-password",
	})

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(credentialsBody))
	s.assertStatusCode(req, 200)
	s.assertStatusCode(req, 400)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(credentialsBody))
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), 200, w.Code)
	assert.True(s.T(), len(w.Result().Cookies()[0].Value) > 2)

	badCredentialsBody1, _ := json.Marshal(auth.Credentials{
		Username: "alex",
		Password: "alex-wrong-password",
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody1))
	s.assertStatusCode(req, 401)

	badCredentialsBody2, _ := json.Marshal(auth.Credentials{
		Username: "shnoo",
		Password: "no-shnoo",
	})
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badCredentialsBody2))
	s.assertStatusCode(req, 401)

	// log in to non existent user
	// check for 401

	// check for token
	// log out the user
	// check for 200
	// log out the user
	// check for 400
	// delete the user
	// check for 200
	// delete the user
	// check for 400
}
