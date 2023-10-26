package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tinyurl/internal/tinyurl"
	"tinyurl/pkg/generator"
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
	r.POST("/make", tinyurl.Make)

	r.GET("/visit", tinyurl.Visit)

	r.GET("/recover", tinyurl.Recover)

	r.Run("0.0.0.0:8089")
	fmt.Println(generator.Hash("https://www.baidu.com/"))
}
