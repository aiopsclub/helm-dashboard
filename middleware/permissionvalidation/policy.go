package permissionvalidation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helm-dashboard/auth"
)

func Policy(policyType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString("username")
		has, err := auth.Enforcer.Enforce(sub, "policy", policyType)
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
				"msg":  fmt.Sprintf("You have not permission to %s policy!", policyType),
			})
			return
		}

	}
}
