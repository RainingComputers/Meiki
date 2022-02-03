package auth_test

import (
	"context"
	"testing"

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
		panic("Why is error notnil")
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
	s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")

	var storedUser auth.User
	s.user_coll.FindOne(s.ctx, bson.M{"username": "shnoo"}).Decode(&storedUser)

	err := bcrypt.CompareHashAndPassword(storedUser.PasswordHash, []byte("thisisveryunsafe"))

	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestShouldCreateTokenForNewUser() {
	s.auth.CreateToken(s.ctx, "alex")
	var storedUserTokens auth.UserTokens
	s.token_coll.FindOne(s.ctx, bson.M{"username": "alex"}).Decode(&storedUserTokens)

	assert.Equal(s.T(), 1, len(storedUserTokens.Tokens))
	assert.True(s.T(), len(storedUserTokens.Tokens[0]) > 2)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
