package account

import (
	"helm-dashboard/internal/encrypthelper"
	"helm-dashboard/internal/jwthelper"
	"helm-dashboard/model"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

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
	has, err := model.Engine.Exist(&model.User{
		Name:   username,
		Passwd: encrypthelper.Md5(password),
	})

	if err != nil {
		c.JSON(401, gin.H{
			"code": "10007",
			"msg":  err,
		})
		return
	}

	if has {
		payload := jwt.Payload{
			Subject:        username,
			ExpirationTime: jwt.NumericDate(time.Now().Add(time.Hour)),
		}
		token, err := jwt.Sign(payload, jwthelper.JwtKey)
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

	} else {
		c.JSON(401, gin.H{
			"code": "10008",
			"msg":  "user not exist!",
		})
	}

}
