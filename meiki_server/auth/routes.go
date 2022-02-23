package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getCreateHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Credentials

		c.BindJSON(&newUser)

		err := a.Create(ctx, newUser.Username, newUser.Password)

		if errors.Is(err, ErrUserAlreadyExists) {
			c.JSON(http.StatusBadRequest, "User already exists")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to create user")
			return
		}
	}
}

func getDeleteHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, "")
	}
}

func getLoginHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, "")
	}
}

func getLogoutHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, "")
	}
}

func CreateRoutes(router *gin.Engine, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
}
