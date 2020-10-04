package account

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var jwtKey = jwt.NewHS256([]byte("helmdashboard"))

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		c.JSON(500, gin.H{
			"code": "10001",
			"msg":  "username and password can't be nil!",
		})
		return
	}

	payload := jwt.Payload{
		Subject:        username,
		ExpirationTime: jwt.NumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.Sign(payload, jwtKey)
	if err != nil {
		c.JSON(401, gin.H{
			"code": "10002",
			"msg":  "token verification failed!",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        "0",
		"msg":         "create token successfully!",
		"accesstoken": token,
	})
}

func subjectFromJWT(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		// Incorrect Authorization header format.
		return ""
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		// JWT not found.
		return ""
	}

	var payload jwt.Payload
	_, err := jwt.Verify([]byte(token), jwtKey, &payload)
	if err != nil {
		return ""
	}
	return payload.Subject
}
