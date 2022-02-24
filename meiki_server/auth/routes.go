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
		var creds Credentials

		c.BindJSON(&creds)

		token, err := a.Login(ctx, creds.Username, creds.Password)

		if err == ErrPasswordMismatch {
			c.JSON(http.StatusUnauthorized, "Password does not match")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to login")
			return
		}

		c.SetCookie("meiki_session_token", string(token), 86400*100, "/", "", true, true)
	}
}

func getLogoutHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		c.BindJSON(&creds)
		token, err := c.Cookie("meiki_session_token")

		if err != nil {
			c.JSON(http.StatusBadRequest, "could not find session token cookie in request")
		}

		err = a.Logout(ctx, creds.Username, []byte(token))

		if err == ErrMissingUserTokens {
			c.JSON(http.StatusBadRequest, "user token is missing")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "unable to logout")
			return
		}
	}
}

func CreateRoutes(router *gin.Engine, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
}
