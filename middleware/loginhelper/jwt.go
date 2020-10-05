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
		c.AbortWithStatusJSON(401, gin.H{"code": 10009, "msg": "Auth Header not found!"})
		return
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		// JWT not found.
		c.AbortWithStatusJSON(401, gin.H{"code": 10008, "msg": "token can't be nil!"})
		return

	}

	var payload jwt.Payload
	_, err := jwt.Verify([]byte(token), jwthelper.JwtKey, &payload)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"code": 10010, "msg": err.Error()})
		return

	}
	c.Set("username", payload.Subject)
	c.Next()

}
