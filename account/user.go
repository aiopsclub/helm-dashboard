package account

import (
	"fmt"
	"helm-dashboard/internal/encrypthelper"
	"helm-dashboard/model"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) == 0 || len(password) == 0 {
		c.JSON(500, gin.H{
			"code": "10001",
			"msg":  "username and password can't be nil!",
		})
		return
	}
	passwdMd5 := encrypthelper.Md5(password)
	user := model.User{}
	user.Name = username
	user.Passwd = passwdMd5
	_, err := model.Engine.Insert(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "10006",
			"msg":  fmt.Sprintf("create user failed! reason: %s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "create user successfully!",
	})
}
