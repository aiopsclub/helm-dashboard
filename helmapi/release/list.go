package release

import (
	"fmt"
	"helm-dashboard/internal/boolhelper"
	"strings"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

type ReleaseListOptions struct {
	Short         bool
	ByDate        bool
	SortReverse   bool
	All           bool
	Uninstalled   bool
	Superseded    bool
	Uninstalling  bool
	Deployed      bool
	Failed        bool
	Pending       bool
	AllNamespaces bool
	Limit         int
	Offset        int
	Filter        string
	Selector      string
}

func List(c *gin.Context) {
	actionConfig := new(action.Configuration)
	listClient := action.NewList(actionConfig)
	listClient.Short = boolhelper.ParseBool(strings.ToLower(c.DefaultQuery("short", "false")))
	fmt.Println(listClient)
	// return listClient.Run()
	results := make([]*release.Release, 0)
	c.JSON(200, results)
}
