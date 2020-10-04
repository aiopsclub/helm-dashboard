package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"helm-dashboard/account"
	_ "helm-dashboard/auth"
	"helm-dashboard/helmapi/release"
	_ "helm-dashboard/model"
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
	v1 := router.Group("/api/v1")

	// v1Release handlers
	v1Release := v1.Group("/release")
	v1Release.GET("/list", release.List)

	// v1Account handlers
	v1Account := v1.Group("/account")
	v1Account.POST("/login", account.Login)

	// v1User handlers
	v1User := v1.Group("/user")
	v1User.POST("/create", account.Create)

	endless.ListenAndServe(":80", router)
}
