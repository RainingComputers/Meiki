package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionCredentials struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func getCreateHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()

		var newUser Credentials

		c.BindJSON(&newUser)

		err := a.Create(ctx, newUser.Username, newUser.Password)

		if errors.Is(err, ErrUserAlreadyExists) {
			c.JSON(http.StatusBadRequest, "Username already exists")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to create user, please try again later")
			return
		}

		c.JSON(http.StatusOK, "Username created")
	}
}

func getDeleteHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		c.BindJSON(&creds)

		match, err := a.PasswordMatches(ctx, creds.Username, creds.Password)

		if err == ErrMissingUser {
			c.JSON(http.StatusBadRequest, "User does not exist")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to check password, please try again later")
			return
		}

		if !match {
			c.JSON(http.StatusUnauthorized, "Password does not match")
			return
		}

		if err := a.Delete(ctx, creds.Username); err != nil {
			c.JSON(http.StatusBadRequest, "Unable to delete user, please try again later")
		}

		c.JSON(http.StatusOK, "Deleted user")
	}
}

func getLoginHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials

		c.BindJSON(&creds)

		token, err := a.Login(ctx, creds.Username, creds.Password)

		if err == ErrMissingUser {
			c.JSON(http.StatusUnauthorized, "User does not exist")
			return
		}

		if err == ErrPasswordMismatch {
			c.JSON(http.StatusUnauthorized, "Password does not match")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to login, please try again later")
			return
		}

		sessionCredentials := SessionCredentials{
			Username: creds.Username,
			Token:    string(token),
		}

		c.JSON(http.StatusOK, sessionCredentials)
	}
}

func getLogoutHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, "User not logged in")
			return
		}

		username := c.GetHeader("X-Username")

		if username == "" {
			c.JSON(http.StatusUnauthorized, "User not logged in")
			return
		}

		err := a.Logout(ctx, username, []byte(token))

		if err == ErrMissingUserTokens {
			fmt.Println(token)
			c.JSON(http.StatusBadRequest, "User token is missing")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "unable to logout")
			return
		}

		c.JSON(http.StatusOK, "Logged out successfully")
	}
}

func getAuthStatus(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, "User not logged in")
			return
		}

		username := c.GetHeader("X-Username")

		if username == "" {
			c.JSON(http.StatusUnauthorized, "User not logged in")
			return
		}

		loggedIn, err := a.Authenticate(ctx, username, []byte(token))

		if errors.Is(err, ErrMissingUserTokens) {
			c.JSON(http.StatusUnauthorized, "No tokens were found for the user")
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unable to authenticate, please try again later")
			return
		}

		if !loggedIn {
			c.JSON(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		c.JSON(http.StatusOK, username)
	}
}

func CreateRoutes(router *gin.Engine, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
	router.GET("/authStatus", getAuthStatus(ctx, auth))
}
