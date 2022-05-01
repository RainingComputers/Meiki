package auth

import (
	"context"
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

const MSG_INVALID_USERNAME = "Username should not contain any special characters other than '-' and '_'"
const MSG_INVALID_PASSWORD = "Password should have minimum five characters"
const MSG_USER_EXISTS = "User already exists"
const MSG_UNABLE_TO_CREATE_USER = "Unable to create user, please try again later"
const MSG_USER_CREATED = "User successfully created"
const MSG_USER_DOES_NOT_EXIST = "User does not exist"
const MSG_UNABLE_TO_DELETE_USER = "Unable to delete user, please try again later"
const MSG_PASSWORD_DOES_NOT_MATCH = "Password does not match"
const MSG_USER_DELETED = "User deleted user successfully"
const MSG_UNABLE_TO_LOGIN = "Unable to login, please try again later"
const MSG_USER_NOT_LOGGED_IN = "User not logged in"
const MSG_TOKEN_DOES_NOT_EXIST = "User token does not exist"
const MSG_UNABLE_TO_LOGOUT = "Unable to logout, please try again later"
const MSG_USER_LOGGED_OUT = "User logged out successfully"
const MSG_UNABLE_TO_AUTHENTICATE = "Unable to authenticate, please try again later"
const MSG_INVALID_OR_WRONG_CREDENTIALS = "Invalid or wrong credentials"
const MSG_UNABLE_TO_PARSE_CREDENTIALS = "Unable to parse credentials in request body json"

func getCreateHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()

		var newUser Credentials
		err := c.BindJSON(&newUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_UNABLE_TO_PARSE_CREDENTIALS)
			return
		}

		err = a.Create(ctx, newUser.Username, newUser.Password)

		if err == ErrInvalidUsername {
			c.JSON(http.StatusBadRequest, MSG_INVALID_USERNAME)
			return
		}

		if err == ErrInvalidPassword {
			c.JSON(http.StatusBadRequest, MSG_INVALID_PASSWORD)
			return
		}

		if err == ErrUserAlreadyExists {
			c.JSON(http.StatusBadRequest, MSG_USER_EXISTS)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_CREATE_USER)
			return
		}

		c.JSON(http.StatusOK, MSG_USER_CREATED)
	}
}

func getDeleteHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		err := c.BindJSON(&creds)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_UNABLE_TO_PARSE_CREDENTIALS)
			return
		}

		match, err := a.PasswordMatches(ctx, creds.Username, creds.Password)

		if err == ErrMissingUser {
			c.JSON(http.StatusBadRequest, MSG_USER_DOES_NOT_EXIST)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_DELETE_USER)
			return
		}

		if !match {
			c.JSON(http.StatusUnauthorized, MSG_PASSWORD_DOES_NOT_MATCH)
			return
		}

		err = a.Delete(ctx, creds.Username)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_UNABLE_TO_CREATE_USER)
			return
		}

		c.JSON(http.StatusOK, MSG_USER_DELETED)
	}
}

func getLoginHandler(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials
		err := c.BindJSON(&creds)

		if err != nil {
			c.JSON(http.StatusBadRequest, MSG_UNABLE_TO_PARSE_CREDENTIALS)
			return
		}

		token, err := a.Login(ctx, creds.Username, creds.Password)

		if err == ErrInvalidUsername {
			c.JSON(http.StatusBadRequest, MSG_INVALID_USERNAME)
			return
		}

		if err == ErrInvalidPassword {
			c.JSON(http.StatusBadRequest, MSG_INVALID_PASSWORD)
			return
		}

		if err == ErrMissingUser {
			c.JSON(http.StatusUnauthorized, MSG_USER_DOES_NOT_EXIST)
			return
		}

		if err == ErrPasswordMismatch {
			c.JSON(http.StatusUnauthorized, MSG_PASSWORD_DOES_NOT_MATCH)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_LOGIN)
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
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			return
		}

		username := c.GetHeader("X-Username")

		if username == "" {
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			return
		}

		err := a.Logout(ctx, username, []byte(token))

		if err == ErrMissingUserTokens {
			fmt.Println(token)
			c.JSON(http.StatusBadRequest, MSG_TOKEN_DOES_NOT_EXIST)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_LOGOUT)
			return
		}

		c.JSON(http.StatusOK, MSG_USER_LOGGED_OUT)
	}
}

func getAuthStatus(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			return
		}

		username := c.GetHeader("X-Username")

		if username == "" {
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			return
		}

		loggedIn, err := a.Authenticate(ctx, username, []byte(token))

		if err == ErrMissingUserTokens {
			c.JSON(http.StatusUnauthorized, MSG_TOKEN_DOES_NOT_EXIST)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_AUTHENTICATE)
			return
		}

		if !loggedIn {
			c.JSON(http.StatusUnauthorized, MSG_INVALID_OR_WRONG_CREDENTIALS)
			return
		}

		c.JSON(http.StatusOK, username)
	}
}

func CreateRoutes(router *gin.RouterGroup, ctx context.Context, auth Auth) {
	router.POST("/create", getCreateHandler(ctx, auth))
	router.DELETE("/delete", getDeleteHandler(ctx, auth))
	router.POST("/login", getLoginHandler(ctx, auth))
	router.POST("/logout", getLogoutHandler(ctx, auth))
	router.GET("/authStatus", getAuthStatus(ctx, auth))
}
