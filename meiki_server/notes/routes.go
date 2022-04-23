package notes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCreateHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func getListHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func getReadHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func getUpdateHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func getDeleteHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func CreateRoutes(router *gin.RouterGroup, ctx context.Context, ns NotesStore) {
	router.POST("/create", getCreateHandler(ctx, ns))
	router.GET("/list", getListHandler(ctx, ns))
	router.GET("/read", getReadHandler(ctx, ns))
	router.POST("/update", getUpdateHandler(ctx, ns))
	router.POST("/delete/:id", getDeleteHandler(ctx, ns))
}
