package tinyurl

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"tinyurl/internal/analyst"
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
	code := c.Param("code")
	if code != "" {
		// get origin from redis cache
		originUrl, err := redis.GetRedisClient().Get("tinyurl:map:" + code).Result()
		if err == nil {
			if originUrl == "" {
				var result bson.M
				mongoClient := mongo.GetMongoClient()
				err = mongoClient.
					Database("tinyurl").
					Collection("url_map").
					FindOne(context.Background(), bson.D{{"short_url", code}}).Decode(&result)
				defer func(mongoClient *mongo2.Client, ctx context.Context) {
					err := mongoClient.Disconnect(ctx)
					if err != nil {
						logrus.WithError(err)
					}
				}(mongoClient, context.Background())
				if err != nil {
					logrus.WithError(err)
				}
				// 将查询结果转换为字符串
				resultStr := fmt.Sprintf("%v", result)
				// 打印或返回结果字符串
				c.Redirect(302, resultStr)
			} else {
				c.Redirect(302, originUrl)
			}
			go analyst.Analysis(code, originUrl, c)
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "not found"})
		}
	}
	return
}

// Recover tinyurl
func Recover(c *gin.Context) {
}
