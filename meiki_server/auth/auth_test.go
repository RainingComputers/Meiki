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
	cancel     context.CancelFunc
	token_coll *mongo.Collection
	user_coll  *mongo.Collection
}

func (s *AuthTestSuite) clean() {
	s.token_coll.DeleteMany(s.ctx, bson.M{})
	s.user_coll.DeleteMany(s.ctx, bson.M{})
}

func (s *AuthTestSuite) SetupTest() {
	log.Initialize()

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)

	client, err := mongo.Connect(s.ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	auth_db := client.Database("auth")
	s.user_coll = auth_db.Collection("users")
	s.token_coll = auth_db.Collection("tokens")

	s.auth, err = auth.CreateAuth(s.ctx, s.token_coll, s.user_coll)
	assert.Nil(s.T(), err)
	s.clean()
}

func (s *AuthTestSuite) TearDownTest() {
	s.clean()
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) TestShouldCreateUser() {
	err := s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.Nil(s.T(), err)

	var storedUser auth.User
	s.user_coll.FindOne(s.ctx, bson.M{"username": "shnoo"}).Decode(&storedUser)

	err = bcrypt.CompareHashAndPassword(storedUser.PasswordHash, []byte("thisisveryunsafe"))

	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestCreateShouldError() {
	err := s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.Nil(s.T(), err)

	err = s.auth.Create(s.ctx, "shnoo", "thisisveryunsafe")
	assert.ErrorIs(s.T(), err, auth.ErrUserAlreadyExists)

	s.cancel()
	err = s.auth.Create(s.ctx, "alex", "alex-password")
	assert.ErrorIs(s.T(), err, context.Canceled)
}

func (s *AuthTestSuite) TestShouldDeleteUser() {
	err := s.auth.Delete(s.ctx, "shnoo")
	assert.ErrorIs(s.T(), err, auth.ErrMissingUser)

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

func (s *AuthTestSuite) TestDeleteShouldError() {
	err := s.auth.Delete(s.ctx, "shnoo")
	assert.ErrorIs(s.T(), err, auth.ErrMissingUser)

	s.cancel()
	err = s.auth.Delete(s.ctx, "shnoo")
	assert.ErrorIs(s.T(), err, context.Canceled)

	// TODO: Simulate error on line 148 using mocks or write internal tests
	// s.token_coll.Drop(s.ctx)
	// err = s.auth.Delete(s.ctx, "shnoo")
	// assert.ErrorIs(s.T(), err, )
}

func (s *AuthTestSuite) TestShouldMatchPassword() {
	// TODO: Errors

	err := s.auth.Create(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err)

	assert.True(s.T(), s.auth.PasswordMatches(s.ctx, "shnoo", "right-password"))
	assert.False(s.T(), s.auth.PasswordMatches(s.ctx, "shnoo", "wrong-password"))
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

// TODO: CreateToken errors?

func (s *AuthTestSuite) TestShouldLogin() {
	err1 := s.auth.Create(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err1)

	token, err2 := s.auth.Login(s.ctx, "shnoo", "right-password")
	assert.Nil(s.T(), err2)
	assert.True(s.T(), len(token) > 0)

	// TODO: Test if tokens are created in DB?

	_, err3 := s.auth.Login(s.ctx, "shnoo", "wrong-password")
	assert.ErrorIs(s.T(), err3, auth.ErrPasswordMismatch)
}

// TODO: Login errors?

// TODO: Test for ReadTokensFromDB?

// TODO: Test for authenticate

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

func (s *AuthTestSuite) TestLogoutShouldError() {
	err := s.auth.Create(s.ctx, "alex", "alex-password")
	assert.Nil(s.T(), err)

	token1, _ := s.auth.CreateToken(s.ctx, "alex")

	err = s.auth.Logout(s.ctx, "doesNotExist", token1)
	assert.ErrorIs(s.T(), err, auth.ErrMissingUserTokens)

	s.cancel()
	err = s.auth.Logout(s.ctx, "alex", token1)
	assert.ErrorIs(s.T(), err, context.Canceled)
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

	_, err = auth.CreateAuth(ctx, token_coll, user_coll)
	assert.ErrorIs(t, err, context.Canceled)
}
