package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getCreateHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()

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

		c.JSON(http.StatusOK, "Username created")
	}
}

func getDeleteHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		c.BindJSON(&creds)

		if !a.PasswordMatches(ctx, creds.Username, creds.Password) {

			// TODO: if the user doesn't exist, return appropriate error OR return bad request
			// We have a toDo in getPasswordHashFromDB that should force us to fix this
			c.JSON(http.StatusBadRequest, "password mismatch OR user doesn't exist while deleting user")
			return
		}

		if err := a.Delete(ctx, creds.Username); err != nil {
			c.JSON(http.StatusBadRequest, "unable to delete user")
		}

		c.JSON(http.StatusOK, "Deleted user")
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

		c.JSON(http.StatusOK, "Logged in successfully")
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

		c.JSON(http.StatusOK, "Logged out successfully")
	}
}

func CreateRoutes(router *gin.Engine, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
}
