package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"time"
	"tinyurl/internal/tinyurl"
)

func main() {
	//rc := redis.GetRedisClient()
	//err := rc.Set("sb", 111, 0).Err()
	//if err != nil {
	//	fmt.Println("Error setting key:", err)
	//} else {
	//	fmt.Println("Key set successfully.")
	//}

	r := gin.Default()

	logger := logrus.New()

	// 设置日志输出，可以同时输出到文件和控制台
	file, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err == nil {
		logger.Out = file
	} else {
		fmt.Println("Failed to log to file, using default stderr")
	}

	r.Use(Logger(logger))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to tinyurl")
	})
	r.POST("/make", tinyurl.Make)

	r.GET("/visit/:code", tinyurl.Visit)

	r.GET("/recover/:shortUrl", tinyurl.Recover)

	err = r.Run("0.0.0.0:8089")
	if err != nil {
		log.Fatalf("gin start error: %v", err)
	}
	//fmt.Println(generator.Hash("https://www.baidu.com/"))
}

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里实现用户识别和计数逻辑
		// 您可以使用浏览器 cookie 或 IP 地址来识别不同用户
		// 然后将统计信息记录到日志文件或数据库

		// 开始时间
		startTime := time.Now()

		c.Next()

		// 结束时间
		endTime := time.Now()

		// 计算请求处理时间
		latency := endTime.Sub(startTime)

		// 记录请求信息
		logger.WithFields(logrus.Fields{
			"client_ip": c.ClientIP(),
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"status":    c.Writer.Status(),
			"latency":   latency,
		}).Info("Request")

		c.Next()
	}
}
