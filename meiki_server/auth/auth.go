package auth

import (
	"context"
	"errors"

	"github.com/RainingComputers/Meiki/log"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	token_coll *mongo.Collection
	user_coll  *mongo.Collection
}

type User struct {
	Username     string `bson:"username"`
	PasswordHash []byte `bson:"password"`
}

type UserTokens struct {
	Username string   `bson:"username"`
	Tokens   [][]byte `bson:"tokens"`
}

func getToken() []byte {
	return []byte(uuid.NewString())
}

func CreateAuth(token_coll *mongo.Collection, user_coll *mongo.Collection) Auth {
	return Auth{token_coll, user_coll}
}

func (a Auth) storeCredentialsInDB(ctx context.Context, user User) error {

	_, err := a.user_coll.InsertOne(ctx, user)

	if err != nil {
		log.Error("Could not store credentials in DB", zap.Error(err))
	}

	return err
}

func (a Auth) Create(ctx context.Context, username string, password string) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	// log.Info("Total Time %d", zap.String("Time difference", end_time.Sub(start_time).String()))

	if err != nil {
		// log.Error("could not create password hash for user", zap.Error(err))
		log.Fatal("could not create password hash for user", zap.Error(err)) // No idea why this would come
		return err
	}

	user := User{
		username,
		passwordHash,
	}

	err = a.storeCredentialsInDB(ctx, user)

	if err != nil {
		log.Error("could not store credentials while creating user", zap.Error(err))
		return err
	}

	return nil
}

func (a Auth) deleteUserInDB(ctx context.Context, username string) error {
	result, err := a.user_coll.DeleteOne(ctx, bson.M{
		"username": username,
	})
	if result.DeletedCount == 0 {
		log.Error("could not find user in DB")
		return errors.New("unable to find user in DB")
	}
	if err != nil {
		log.Error("unable to delete user from DB", zap.Error(err))
		return err
	}
	return nil
}

func (a Auth) deleteUserTokensInDB(ctx context.Context, username string) error {
	result, err := a.token_coll.DeleteOne(ctx, bson.M{
		"username": username,
	})
	if result.DeletedCount == 0 {
		log.Error("could not find user token in DB")
		return errors.New("unable to find user tokens in DB")
	}
	if err != nil {
		log.Error("unable to delete user token from DB", zap.Error(err))
		return err
	}
	return nil
}

func (a Auth) Delete(ctx context.Context, username string) error {
	err := a.deleteUserInDB(ctx, username)
	if err != nil {
		log.Error("unable to delete user", zap.Error(err))
		return err
	}

	err = a.deleteUserTokensInDB(ctx, username)
	if err != nil {
		log.Error("unable to delete user tokens", zap.Error(err))
		return err
	}
	return nil
}

func (a Auth) CreateToken(ctx context.Context, username string) ([]byte, error) {
	newToken := getToken()

	result, err := a.token_coll.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"tokens": newToken}})

	if err != nil {
		log.Error("Unable to update token into DB", zap.Error(err))
		return nil, err
	}

	if result.MatchedCount == 0 {
		newUserTokens := UserTokens{
			Username: username,
			Tokens:   [][]byte{newToken},
		}

		_, err := a.token_coll.InsertOne(ctx, newUserTokens)

		if err != nil {
			log.Error("Unable to insert token into DB", zap.Error(err))
			return nil, err
		}
	}

	return newToken, nil
}

func (a Auth) Authenticate(username string, token string) (bool, error) {
	// get tokens struct with username

	// check if token exists in token array

	// if yes return true else false

	return false, errors.New("test errors")
}

func (a Auth) Logout(username string, token string) error {
	// delete particular token from array by username

	return errors.New("test errors")
}

func TestLog() {
	log.Error("something", zap.Error(errors.New("some error")))
}
