package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"helm-dashboard/account"
	"helm-dashboard/auth"
	"helm-dashboard/helmapi/release"
	_ "helm-dashboard/model"
	"time"
)

func init() {
	gin.DisableConsoleColor()
}

func configMiddleware(r *gin.Engine) {
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
	v1Release.GET("/", release.List)

	// v1Account handlers
	v1Account := v1.Group("/account")
	v1Account.POST("/login", account.Login)

	// v1User handlers
	v1User := v1.Group("/user")
	v1User.POST("/", account.Create)

	// v1Policy handlers
	v1Policy := v1.Group("/policy")
	v1Policy.POST("/", auth.PolicyAdd)
	v1Policy.GET("/", auth.PolicyList)
	v1Policy.DELETE("/", auth.PolicyRemove)

	endless.ListenAndServe(":80", router)
}
