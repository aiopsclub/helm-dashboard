package auth

import (
	"fmt"
	"os"
	"path"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var Enforcer *casbin.Enforcer

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbPath := path.Join(cwd, "db/system.db")
	fmt.Println(dbPath)

	adapter, err := xormadapter.NewAdapter("sqlite3", dbPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	casbinConfPath := path.Join(cwd, "conf/basic_with_root_model.conf")
	fmt.Println(casbinConfPath)
	Enforcer, err = casbin.NewEnforcer(casbinConfPath, adapter)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Enforcer.LoadPolicy()
}

func PolicyAdd(c *gin.Context) {
	_, err := Enforcer.AddPolicy("a", "b", "c")
	if err != nil {
		c.JSON(500, gin.H{
			"code": "10003",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "policy add successfully",
	})
}

func PolicyRemove(c *gin.Context) {
	_, err := Enforcer.RemovePolicy("a", "b", "c")
	if err != nil {
		c.JSON(500, gin.H{
			"code": "10004",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "policy remove successfully",
	})
}

func PolicyList(c *gin.Context) {
	permissionsforuser := Enforcer.GetPermissionsForUser("a")
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "policy get successfully",
		"data": permissionsforuser,
	})
}
