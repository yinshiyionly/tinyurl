package mongo

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

const (
    URI           = "mongodb://127.0.0.1:27017"
    DATABASE      = "tinyurl"
    TIMEOUT       = 10
    POOL_MAX_SIZE = 10
    USERNAME      = "root"
    PASSWORD      = "example"
)

var mongoClient *mongo.Client

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
    credential := options.Credential{
        Username: USERNAME,
        Password: PASSWORD,
    }
    defer cancel()
    o := options.Client().ApplyURI(URI).SetAuth(credential)
    o.SetMaxPoolSize(POOL_MAX_SIZE)
    var err error
    mongoClient, err = mongo.Connect(ctx, o)
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}

func GetMongoClient() *mongo.Client {
    return mongoClient
}
