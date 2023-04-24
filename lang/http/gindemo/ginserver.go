package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const KeyRequestId = "requestId"

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {

		s := time.Now()

		c.Next() //继续执行

		//path,response code,log latency,
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)

	}, func(c *gin.Context) {

		c.Set(KeyRequestId, rand.Int())

		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {

		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(KeyRequestId); exists {
			h[KeyRequestId] = rid
		}
		c.JSON(http.StatusOK, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
