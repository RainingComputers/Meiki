package auth_test

import (
	"context"
	"testing"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type ValueObject struct {
	Value string
}

func TestSomething(t *testing.T) {

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("Why is error notnil")
	}

	auth_db := client.Database("auth")
	user_coll := auth_db.Collection("users")
	token_coll := auth_db.Collection("tokens")

	a := auth.Auth{
		User_coll: user_coll, Token_coll: token_coll,
	}

	a.Create(ctx, "shnoo", "thisisveryunsafe")

	insertedDoc := user_coll.FindOne(ctx, bson.M{"username": "shnoo"})
	var storedUser auth.User
	_ = insertedDoc.Decode(&storedUser)
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte("thisisveryunsafe"))
	assert.Nil(t, err)
}
