package auth

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	token_coll *mongo.Collection
	user_coll  *mongo.Collection
}

func (a Auth) CreateToken(username string) (string, error) {
	return "random", errors.New("Test errors")
}

func (a Auth) DeleteToken(username string, token string) error {
	return errors.New("Test errors")
}

func (a Auth) Create(username string, password string) error {
	return errors.New("Test errors")
}

func (a Auth) Authenticate(username string, token string) (bool, error) {
	return false, errors.New("Test errors")
}
