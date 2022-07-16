package main

import (
	"context"
	"os"
	"time"

	"github.com/RainingComputers/Meiki/auth"
	"github.com/RainingComputers/Meiki/health"
	"github.com/RainingComputers/Meiki/log"
	"github.com/RainingComputers/Meiki/notes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func run() error {
	databaseName := getEnv("MEIKI_DATABASE_NAME", "meiki")
	databaseURL := getEnv("MEIKI_DATABASE_URL", "mongodb://root:example@localhost:27017")
	corsOriginURL := getEnv("CORS_ORIGIN_URL", "http://localhost:3000")

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))

	if err != nil {
		log.Error("unable to connect to mongo", zap.Error(err))
		return err
	}

	meikiDB := client.Database(databaseName)
	userColl := meikiDB.Collection("users")
	tokenColl := meikiDB.Collection("tokens")
	notesColl := meikiDB.Collection("notes")

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
		AllowOrigins:     []string{corsOriginURL},
		AllowMethods:     []string{"*", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin, Content-Type, Access-Control-Allow-Headers", "X-Username", "X-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	healthRouter := router.Group("/health")
	health.CreateRoutes(healthRouter)

	authRouter := router.Group("/auth")
	auth.CreateRoutes(authRouter, ctx, authController)

	notesRouter := router.Group("/notes")
	notesRouter.Use(auth.GetAuthMiddleware(ctx, authController))
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
	log.Fatal("an error has occurred has in initializing the app", zap.Error(run()))
}
