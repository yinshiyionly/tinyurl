package analyst

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Analysis(code, originUrl string, c *gin.Context) {
	// todo analysis
	log.Fatalln("IP" + c.ClientIP())
	return
}
