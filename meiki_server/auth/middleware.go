package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthMiddleware(ctx context.Context, a Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			c.Abort()
			return
		}

		username := c.GetHeader("X-Username")
		if username == "" {
			c.JSON(http.StatusUnauthorized, MSG_USER_NOT_LOGGED_IN)
			c.Abort()
			return
		}

		authorized, err := a.Authenticate(ctx, username, token)

		if err == ErrMissingUserTokens {
			c.JSON(http.StatusUnauthorized, MSG_TOKEN_DOES_NOT_EXIST)
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, MSG_UNABLE_TO_AUTHENTICATE)
			c.Abort()
			return
		}

		if !authorized {
			c.JSON(http.StatusUnauthorized, MSG_INVALID_OR_WRONG_CREDENTIALS)
			c.Abort()
			return
		}
	}
}
