package main

import (
	"context"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/RainingComputers/Meiki/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func run() error {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	auth_db := client.Database("auth")
	userColl := auth_db.Collection("users")
	tokenColl := auth_db.Collection("tokens")

	authCtx, err := auth.CreateAuth(ctx, tokenColl, userColl)

	if err != nil {
		return err
	}

	router := gin.Default()

	auth.CreateRoutes(router, ctx, authCtx)

	err = router.Run()

	if err != nil {
		return err
	}

	return nil
}

func main() {
	log.Initialize()
	log.Fatal("An error has occurred has in initializing the app", zap.Error(run()))
}
