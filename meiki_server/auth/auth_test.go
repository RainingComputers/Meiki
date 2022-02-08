package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/RainingComputers/Meiki/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type AuthTestSuite struct {
	suite.Suite
	ctx        context.Context
	auth       auth.Auth
	token_coll *mongo.Collection
	user_coll  *mongo.Collection
}

func (s *AuthTestSuite) clean() {
	s.token_coll.DeleteMany(s.ctx, bson.M{})
	s.user_coll.DeleteMany(s.ctx, bson.M{})
}

func (s *AuthTestSuite) SetupTest() {
	log.Initialize()
	s.ctx = context.Background()

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	auth_db := client.Database("auth")
	s.user_coll = auth_db.Collection("users")
	s.token_coll = auth_db.Collection("tokens")

	s.auth = auth.CreateAuth(s.token_coll, s.user_coll)
	s.clean()
}

func (s *AuthTestSuite) TeardownTest() {
	s.clean()
}

func (s *AuthTestSuite) TestShouldCreateUser() {
	err := s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.Nil(s.T(), err)

	var storedUser auth.User
	s.user_coll.FindOne(s.ctx, bson.M{"username": "shnoo"}).Decode(&storedUser)

	err = bcrypt.CompareHashAndPassword(storedUser.PasswordHash, []byte("thisisveryunsafe"))

	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestShouldCreateTokenForNewUser() {
	token, err := s.auth.CreateToken(s.ctx, "alex")
	assert.Nil(s.T(), err)

	var storedUserTokens auth.UserTokens
	s.token_coll.FindOne(s.ctx, bson.M{"username": "alex"}).Decode(&storedUserTokens)

	assert.Equal(s.T(), 1, len(storedUserTokens.Tokens))
	assert.True(s.T(), len(storedUserTokens.Tokens[0]) > 2)
	assert.Equal(s.T(), token, storedUserTokens.Tokens[0])
}

func (s *AuthTestSuite) TestShouldAddTokenForExistingUser() {
	token1, err1 := s.auth.CreateToken(s.ctx, "alex")
	assert.Nil(s.T(), err1)

	token2, err2 := s.auth.CreateToken(s.ctx, "alex")
	assert.Nil(s.T(), err2)

	var storedUserTokens auth.UserTokens
	s.token_coll.FindOne(s.ctx, bson.M{"username": "alex"}).Decode(&storedUserTokens)

	assert.Equal(s.T(), 2, len(storedUserTokens.Tokens))
	assert.True(s.T(), len(storedUserTokens.Tokens[0]) > 2)
	assert.True(s.T(), len(storedUserTokens.Tokens[1]) > 2)
	assert.Equal(s.T(), token1, storedUserTokens.Tokens[0])
	assert.Equal(s.T(), token2, storedUserTokens.Tokens[1])
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func TestShouldErroroutWhenMongoCannotBeConnected(t *testing.T) {
	log.Initialize()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	assert.Nil(t, err)
	auth_db := client.Database("auth")
	user_coll := auth_db.Collection("users")
	token_coll := auth_db.Collection("tokens")

	auth := auth.CreateAuth(token_coll, user_coll)
	err = auth.Create(ctx, "shnoo", "thisisveryunsafe")
	assert.NotNil(t, err)
}
