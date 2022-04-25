package main

import (
	"context"
	"time"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/RainingComputers/Meiki/log"
	"github.com/RainingComputers/Meiki/notes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func run() error {
	// TODO: Make config env variables and use same db for auth and notes

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		panic("unable to connect to mongo for test suite")
	}

	meiki_db := client.Database("meiki")
	userColl := meiki_db.Collection("users")
	tokenColl := meiki_db.Collection("tokens")
	notesColl := meiki_db.Collection("notes")

	authController, err := auth.CreateAuth(ctx, tokenColl, userColl)
	if err != nil {
		return err
	}

	notesStoreController, err := notes.CreateNotesStore(ctx, notesColl)

	if err != nil {
		return err
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin, Content-Type, Access-Control-Allow-Headers", "X-Username", "X-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRouter := router.Group("/auth")
	auth.CreateRoutes(authRouter, ctx, authController)

	notesRouter := router.Group("/notes")
	notes.CreateRoutes(notesRouter, ctx, notesStoreController)

	err = router.Run()

	if err != nil {
		return err
	}

	client.Disconnect(ctx)

	return nil
}

func main() {
	log.Initialize()
	log.Fatal("An error has occurred has in initializing the app", zap.Error(run()))
}
