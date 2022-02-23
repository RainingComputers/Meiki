package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world")
	}
}

func getCreateHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	type CreateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(c *gin.Context) {
		var newUser CreateUserRequest

		c.BindJSON(&newUser)

		err := a.Create(ctx, newUser.Username, newUser.Password)

		if err != nil {
			// TODO: return BAD_REQUEST if the user already exists, for now this
			c.JSON(http.StatusInternalServerError, "Unable to create user")
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
	router.POST("/hello", HelloHandler())
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
}
