package tinyurl

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "log"
    "net/http"
    "tinyurl/internal/mongo"
    "tinyurl/internal/redis"
    "tinyurl/pkg/generator"
)

type makeParams struct {
    Url string `json:"url" binding:"required,max=500"`
}

// Make tinyurl
// POST
func Make(c *gin.Context) {
    var data makeParams
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // make tinyurl
    shortUrl := generator.Hash(data.Url)[0]
    // save mongodb
    collection := mongo.GetMongoClient().Database("tinyurl").Collection("url_map")
    _, err := collection.InsertOne(context.TODO(), bson.D{{"origin_url", data.Url}, {"short_url", shortUrl}})
    if err != nil {
        log.Fatalf("insert to mongodb err: %v", err)
    }
    // save redis
    redis.GetRedisClient().Set("tinyurl:map:"+shortUrl, data.Url, 0)
    // return
    c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "data":    shortUrl,
    })
}

// Visit tinyurl
func Visit(c *gin.Context) {
    shortUrl := c.Param("shortUrl")
    if shortUrl != "" {
        // get origin from redis cache
        originUrl, err := redis.GetRedisClient().Get("tinyurl:map:" + shortUrl).Result()
        if err == nil {
            if originUrl == "" {
                var result bson.M
                mongoClient := mongo.GetMongoClient()
                err = mongoClient.
                    Database("tinyurl").
                    Collection("url_map").
                    FindOne(context.Background(), bson.D{{"short_url", shortUrl}}).Decode(&result)
                defer mongoClient.Disconnect(context.Background())
                if err != nil {
                    log.Fatalf("Error querying MongoDB: %v", err)
                }
                // 将查询结果转换为字符串
                resultStr := fmt.Sprintf("%v", result)
                // 打印或返回结果字符串
                fmt.Println("Query Result:", resultStr)
                c.Redirect(302, resultStr)
                c.JSON(http.StatusOK, gin.H{"message": "Key does not exist"})
            } else {
                c.Redirect(302, originUrl)
            }
        } else {
            c.JSON(http.StatusOK, gin.H{"message": "not found"})
        }
    }
    return
}

// Recover tinyurl
func Recover(c *gin.Context) {

}
