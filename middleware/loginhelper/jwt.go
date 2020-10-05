package loginhelper

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"helm-dashboard/internal/jwthelper"
	"strings"
)

func RequireJWT(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		// Incorrect Authorization header format.
		c.Abort()
		return
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		// JWT not found.
		c.Abort()
		return

	}

	var payload jwt.Payload
	_, err := jwt.Verify([]byte(token), jwthelper.JwtKey, &payload)
	if err != nil {
		c.Abort()
		return

	}
	c.Next()

}
