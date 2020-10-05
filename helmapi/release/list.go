package release

import (
	"fmt"
	"helm-dashboard/helmapi/initialize"
	"helm-dashboard/internal/boolhelper"
	"strings"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func List(c *gin.Context) {
	results := make([]*release.Release, 0)
	listClient := action.NewList(initialize.ActionConfig)
	listClient.Short = boolhelper.ParseBool(strings.ToLower(c.DefaultQuery("short", "false")))
	releases, err := listClient.Run()
	results = append(results, releases...)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 10011,
			"msg":  fmt.Sprintf("get release list failed! reason: %s", err),
		})
		return
	}
	fmt.Println(results)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "get release list successfully!",
		"data": results,
	})

}
