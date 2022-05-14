package auth

import (
	"context"
	"net/http"

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
			c.JSON(http.StatusBadRequest, MSG_UNABLE_TO_DELETE_USER)
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
		username := c.GetHeader("X-Username")
		err := a.Logout(ctx, username, token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_LOGOUT)
			return
		}

		c.JSON(http.StatusOK, MSG_USER_LOGGED_OUT)
	}
}

func getAuthStatus(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Authorized")
	}
}

func CreateRoutes(router *gin.RouterGroup, ctx context.Context, auth Auth) {
	unauthorizedRouter := router.Group("/")
	unauthorizedRouter.POST("/create", getCreateHandler(ctx, auth))
	unauthorizedRouter.DELETE("/delete", getDeleteHandler(ctx, auth))
	unauthorizedRouter.POST("/login", getLoginHandler(ctx, auth))

	authorizedRouter := router.Group("/")
	authorizedRouter.Use(GetAuthMiddleware(ctx, auth))
	authorizedRouter.POST("/logout", getLogoutHandler(ctx, auth))
	authorizedRouter.GET("/authStatus", getAuthStatus(ctx, auth))
}
