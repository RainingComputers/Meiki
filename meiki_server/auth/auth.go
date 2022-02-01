package auth

import (
	"context"
	"errors"
	"log"

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
	PasswordHash string `bson:"password"`
}

func (a Auth) CreateToken(username string) (string, error) {
	return "random", errors.New("Test errors")
}

func (a Auth) DeleteToken(username string, token string) error {
	return errors.New("Test errors")
}

func getToken() string {
	return uuid.NewString()
}

func (a Auth) storeCredentialsInDB(ctx context.Context, user User) error {

	_, err := a.User_coll.InsertOne(ctx, user)

	if err != nil {
		panic(err)
	}

	// _, err = coll.DeleteMany(ctx, bson.M{})

	// if err != nil {
	// 	panic(err)
	return nil
}

func (a Auth) Create(ctx context.Context, username string, password string) error {

	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic("could not create passwordHash")
	}

	passwordHash := string(passwordHashByte)

	user := User{
		username,
		passwordHash,
	}

	err = a.storeCredentialsInDB(ctx, user)

	if err != nil {
		log.Panic("could not save credentials")
	}

	return errors.New("Test errors")
}

func (a Auth) Authenticate(username string, token string) (bool, error) {
	return false, errors.New("Test errors")
}
