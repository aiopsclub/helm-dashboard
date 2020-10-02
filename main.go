package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helm-dashboard/helmapi"
	"time"
)

func init() {
	gin.DisableConsoleColor()
	helmapi.Test()
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
	r := gin.Default()
	configMiddleware(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
