package notes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const MSG_UNABLE_TO_CREATE_NOTE = "Unable to create note, please try again later"
const MSG_PARSE_ERROR = "Unable to parse request body"
const MSG_INVALID_TITLE = "Invalid note title"
const MSG_NOTE_DOES_NOT_EXIST = "Note does not exist"
const MSG_INVALID_ID = "Invalid id"
const MSG_UNABLE_TO_PERFORM_ACTION = "Unable perform requested action"

type TitleRequest struct {
	Title string `json:"title"`
}

type UpdateRequest struct {
	Content string `json:"content"`
}

func errorToResponse(c *gin.Context, err error) {
	if err == ErrNoteDoesNotExist {
		c.JSON(http.StatusBadRequest, MSG_NOTE_DOES_NOT_EXIST)
		return
	}

	if err == ErrInvalidId {
		c.JSON(http.StatusBadRequest, MSG_INVALID_ID)
		return
	}

	if err == ErrInvalidTitle {
		c.JSON(http.StatusBadRequest, MSG_INVALID_TITLE)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_PERFORM_ACTION)
		return
	}
}

func getCreateHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createRequest TitleRequest
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

		id, err := ns.Create(ctx, note)

		if err != nil {
			errorToResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, id)
	}
}

func getListHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")

		notesResponseList, err := ns.List(ctx, username)

		if err != nil {
			errorToResponse(c, err)
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

		if err != nil {
			errorToResponse(c, err)
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

		err = ns.Update(ctx, id, username, updateRequest.Content)

		if err != nil {
			errorToResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, "Updated note")
	}
}

func genRenameHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var renameRequest TitleRequest
		err := c.BindJSON(&renameRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_PARSE_ERROR)
			return
		}

		id := c.Param("id")
		username := c.GetHeader("X-Username")

		err = ns.Rename(ctx, id, username, renameRequest.Title)

		if err != nil {
			errorToResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, "Renamed note")
	}
}

func getDeleteHandler(ctx context.Context, ns NotesStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		username := c.GetHeader("X-Username")

		err := ns.Delete(ctx, id, username)

		if err != nil {
			errorToResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, "Deleted note")
	}
}

func CreateRoutes(router *gin.RouterGroup, ctx context.Context, ns NotesStore) {
	router.POST("/create", getCreateHandler(ctx, ns))
	router.GET("/list", getListHandler(ctx, ns))
	router.GET("/read/:id", getReadHandler(ctx, ns))
	router.PUT("/update/:id", getUpdateHandler(ctx, ns))
	router.PUT("/rename/:id", genRenameHandler(ctx, ns))
	router.DELETE("/delete/:id", getDeleteHandler(ctx, ns))
}
