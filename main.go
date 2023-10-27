package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
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
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to tinyurl")
    })
    r.POST("/make", tinyurl.Make)

    r.GET("/visit/:shortUrl", tinyurl.Visit)

    r.GET("/recover", tinyurl.Recover)

    err := r.Run("0.0.0.0:8089")
    if err != nil {
        log.Fatalf("gin start error: %v", err)
    }
    //fmt.Println(generator.Hash("https://www.baidu.com/"))
}
