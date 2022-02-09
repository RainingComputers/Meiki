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

	s.auth = auth.CreateAuth(s.ctx, s.token_coll, s.user_coll)
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

func (s *AuthTestSuite) TestShouldErrorOutIfUserExists() {
	err := s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.Nil(s.T(), err)

	err = s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.NotNil(s.T(), err)
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
	token1, _ := s.auth.CreateToken(s.ctx, "alex")

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

func (s *AuthTestSuite) TestShouldDeleteUser() {
	err := s.auth.Delete(s.ctx, "shnoo")
	assert.NotNil(s.T(), err)
	// assert.ErrorIsf(s.T(), err, errors.ErrKeyIncorrect, "unable to find user in DB")

	err = s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.Nil(s.T(), err)

	_, err = s.auth.CreateToken(s.ctx, "shnoo")
	assert.Nil(s.T(), err)

	err = s.auth.Delete(s.ctx, "shnoo")
	assert.Nil(s.T(), err)

	var storedUser auth.User
	s.user_coll.FindOne(s.ctx, bson.M{"username": "shnoo"}).Decode(&storedUser)

	assert.Equal(s.T(), 0, len(storedUser.Username))
	assert.Equal(s.T(), 0, len(storedUser.PasswordHash))

	var storedUserToken auth.UserTokens
	s.token_coll.FindOne(s.ctx, bson.M{"username": "shnoo"}).Decode(&storedUserToken)

	assert.Equal(s.T(), 0, len(storedUserToken.Tokens))
	assert.Equal(s.T(), 0, len(storedUserToken.Username))

}

func (s *AuthTestSuite) TestShouldMatchPassword() {
	err := s.auth.Create(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err)

	assert.True(s.T(), s.auth.PasswordMatches(s.ctx, "shnoo", "right-password"))
	assert.False(s.T(), s.auth.PasswordMatches(s.ctx, "shnoo", "wrong-password"))
}

func (s *AuthTestSuite) TestShouldLogin() {
	err1 := s.auth.Create(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err1)

	token, err2 := s.auth.Login(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err2)

	assert.True(s.T(), len(token) > 0)
}

func (s *AuthTestSuite) TestShouldLogout() {
	err := s.auth.Create(s.ctx, "alex", "alex-password")
	assert.Nil(s.T(), err)

	token1, _ := s.auth.CreateToken(s.ctx, "alex")
	token2, _ := s.auth.CreateToken(s.ctx, "alex")

	s.auth.Logout(s.ctx, "alex", token1)

	storedTokens := s.auth.ReadTokensFromDB(s.ctx, "alex")
	assert.Equal(s.T(), len(storedTokens), 1)
	assert.Equal(s.T(), storedTokens[0], token2)

	s.auth.Logout(s.ctx, "alex", token2)

	storedTokens = s.auth.ReadTokensFromDB(s.ctx, "alex")
	assert.Equal(s.T(), len(storedTokens), 0)
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

	auth := auth.CreateAuth(ctx, token_coll, user_coll)
	err = auth.Create(ctx, "shnoo", "thisisveryunsafe")
	assert.NotNil(t, err)

	// Need to check the correct error for future

}
