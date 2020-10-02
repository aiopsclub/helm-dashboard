package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"helm-dashboard/helmapi/release"
	"time"
)

func init() {
	gin.DisableConsoleColor()
}

func configMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

	}))
}

func main() {
	router := gin.Default()
	configMiddleware(router)
	v1 := router.Group("/v1")

	// v1Release handlers
	v1Release := v1.Group("/release")
	v1Release.GET("/list", release.List)

	endless.ListenAndServe(":80", router)
}
