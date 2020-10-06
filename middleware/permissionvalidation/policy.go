package permissionvalidation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helm-dashboard/auth"
)

func Policy(objectType, policyType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString("username")

		switch objectType {
		case "policy":
		case "release":
			objectType = c.Query("releasename")
			if len(objectType) == 0 {
				c.AbortWithStatusJSON(500, gin.H{
					"code": 10015,
					"msg":  "releasename can't be nil!",
				})
				return
			}
		default:
		}
		has, err := auth.Enforcer.Enforce(sub, objectType, policyType)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"code": 10013,
				"msg":  err.Error(),
			})
			return
		}

		if has {
			c.Next()
		} else {
			c.AbortWithStatusJSON(403, gin.H{
				"code": 10012,
				"msg":  fmt.Sprintf("You have not permission to %s %s!", policyType, objectType),
			})
			return
		}

	}
}
