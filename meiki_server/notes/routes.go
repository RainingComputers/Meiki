package notes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const MSG_UNABLE_TO_CREATE_NOTE = "Unable to create note"
const MSG_PARSE_ERROR = "Unable to parse request body"

type CreateRequest struct {
	Title string `json:"title"`
}

type UpdateRequest struct {
	Content string `json:"content"`
}

func getCreateHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createRequest CreateRequest
		err := c.BindJSON(&createRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_PARSE_ERROR)
			return
		}

		username := c.GetHeader("X-Username")

		note := Note{
			Username: username,
			Title:    createRequest.Title,
			Content:  "",
		}

		noteResponse, err := ns.Create(ctx, note)

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_CREATE_NOTE)
			return
		}

		c.JSON(http.StatusOK, noteResponse)
	}
}

func getListHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")

		notesResponseList, err := ns.List(ctx, username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to list notes")
			return
		}

		c.JSON(http.StatusOK, notesResponseList)
	}
}

func getReadHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		username := c.GetHeader("X-Username")

		content, err := ns.Read(ctx, id, username)

		if err == ErrNoteDoesNotExist {
			c.JSON(http.StatusBadRequest, "Note does not exist")
			return
		}

		if err == ErrInvalidId {
			c.JSON(http.StatusBadRequest, "Invalid id")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to delete note")
			return
		}

		c.JSON(http.StatusOK, content)
	}
}

func getUpdateHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateRequest UpdateRequest
		err := c.BindJSON(&updateRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_PARSE_ERROR)
			return
		}

		id := c.Param("id")
		username := c.GetHeader("X-Username")

		err = ns.Update(ctx, id, updateRequest.Content, username)

		if err == ErrNoteDoesNotExist {
			c.JSON(http.StatusBadRequest, "Note does not exist") // TODO: DRY these if statements
			return
		}

		if err == ErrInvalidId {
			c.JSON(http.StatusBadRequest, "Invalid id")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to delete note")
			return
		}

		c.JSON(http.StatusOK, "Updated note")
	}
}

func getDeleteHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		username := c.GetHeader("X-Username")

		err := ns.Delete(ctx, id, username)

		if err == ErrNoteDoesNotExist {
			c.JSON(http.StatusBadRequest, "Note does not exist")
			return
		}

		if err == ErrInvalidId {
			c.JSON(http.StatusBadRequest, "Invalid id")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to delete note")
			return
		}

		c.JSON(http.StatusOK, "Deleted note")
	}
}

func CreateRoutes(router *gin.RouterGroup, ctx context.Context, ns NotesStore) {
	router.POST("/create", getCreateHandler(ctx, ns))
	router.GET("/list", getListHandler(ctx, ns))
	router.GET("/read/:id", getReadHandler(ctx, ns))
	router.POST("/update/:id", getUpdateHandler(ctx, ns))
	router.POST("/delete/:id", getDeleteHandler(ctx, ns))
}
