package auth

import (
	"context"
	"errors"

	"github.com/RainingComputers/Meiki/log"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Token_coll *mongo.Collection
	User_coll  *mongo.Collection
}

type User struct {
	Username     string `bson:"username"`
	PasswordHash []byte `bson:"password"`
}

type UserTokens struct {
	Username string `bson:"username"`
	Tokens   []string
}

func getToken() string {
	return uuid.NewString()
}

func (a Auth) storeCredentialsInDB(ctx context.Context, user User) error {

	_, err := a.User_coll.InsertOne(ctx, user)

	if err != nil {
		log.Error("Could not store credentials in DB", zap.Error(err))
	}

	return err
}

func (a Auth) Create(ctx context.Context, username string, password string) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Error("Could not create password hash for user", zap.Error(err))
		return err
	}

	user := User{
		username,
		passwordHash,
	}

	err = a.storeCredentialsInDB(ctx, user)

	if err != nil {
		log.Error("Could not store credentials while creating user", zap.Error(err))
		return err
	}

	return errors.New("Test errors")
}

func (a Auth) CreateToken(username string) (string, error) {
	// check if usertoken struct exists in mongo

	// if not create one

	// insert this token into token array

	return "random", errors.New("Test errors")
}

func (a Auth) Authenticate(username string, token string) (bool, error) {
	// get tokens struct with usernmae

	// check if token exists in token array

	// if yes return true else false

	return false, errors.New("Test errors")
}

func (a Auth) Logout(username string, token string) error {
	// delete particular token from array by username

	return errors.New("Test errors")
}

func TestLog() {
	log.Error("something", zap.Error(errors.New("some error")))
}
