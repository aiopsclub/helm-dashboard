package history

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helm-dashboard/helmapi/initialize"
	"helm-dashboard/internal/commonhelper"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
	helmtime "helm.sh/helm/v3/pkg/time"
	"strconv"
)

type releaseInfo struct {
	Revision    int           `json:"revision"`
	Updated     helmtime.Time `json:"updated"`
	Status      string        `json:"status"`
	Chart       string        `json:"chart"`
	AppVersion  string        `json:"app_version"`
	Description string        `json:"description"`
}

type releaseHistory []releaseInfo

func formatChartname(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return fmt.Sprintf("%s-%s", c.Name(), c.Metadata.Version)
}

func formatAppVersion(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return c.AppVersion()
}

func getReleaseHistory(rls []*release.Release) (history releaseHistory) {
	for i := len(rls) - 1; i >= 0; i-- {
		r := rls[i]
		c := formatChartname(r.Chart)
		s := r.Info.Status.String()
		v := r.Version
		d := r.Info.Description
		a := formatAppVersion(r.Chart)

		rInfo := releaseInfo{
			Revision:    v,
			Status:      s,
			Chart:       c,
			AppVersion:  a,
			Description: d,
		}
		if !r.Info.LastDeployed.IsZero() {
			rInfo.Updated = r.Info.LastDeployed

		}
		history = append(history, rInfo)
	}

	return history
}

func List(c *gin.Context) {
	listClient := action.NewHistory(initialize.ActionConfig)
	releaseName := c.Query("releasename")
	max, err := strconv.Atoi(c.DefaultQuery("max", "256"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": 10016,
			"msg":  "max arg mush be int string!",
		})
		return
	}
	listClient.Max = max

	hist, err := listClient.Run(releaseName)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 10014,
			"msg":  fmt.Sprintf("get release history failed! reason: %s", err),
		})
		return
	}

	releaseutil.Reverse(hist, releaseutil.SortByRevision)

	var rels []*release.Release
	for i := 0; i < commonhelper.Min(len(hist), listClient.Max); i++ {
		rels = append(rels, hist[i])
	}

	if len(rels) == 0 {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "get release history successfully!",
			"data": releaseHistory{},
		})
		return
	}

	releaseHistory := getReleaseHistory(rels)

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "get release history successfully!",
		"data": releaseHistory,
	})

}
