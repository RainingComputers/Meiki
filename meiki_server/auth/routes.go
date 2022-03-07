package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/RainingComputers/Meiki/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

		c.SetCookie("meiki_session_token", string(token), 86400*100, "/", "http://localhost", true, true)
		c.SetCookie("meiki_username", creds.Username, 86400*100, "/", "http://localhost", true, true)
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

func getAuthStatus(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("meiki_session_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, "not logged in. Redirecting...")
			return
		}
		username, err := c.Cookie("meiki_username")

		if err != nil {
			c.JSON(http.StatusUnauthorized, "not logged in. Redirecting...")
			return
		}

		loggedIn, err := a.Authenticate(ctx, username, []byte(sessionToken))

		if err != nil {
			log.Error("Unable to create unique index in token collection", zap.Error(err))
			c.JSON(http.StatusInternalServerError, "unable to authenticate")
			return
		}

		if !loggedIn {
			c.JSON(http.StatusUnauthorized, "invalid credientials")
			return
		}

		c.JSON(http.StatusOK, "Login Successful")

	}
}

func CreateRoutes(router *gin.Engine, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.POST("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
	router.POST("/authStatus", getAuthStatus(ctx, auth))
}
