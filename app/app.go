package app

import (
	"example.com/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map the urls", "step:1", "status:pending")
	mapUrls()
	log.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
