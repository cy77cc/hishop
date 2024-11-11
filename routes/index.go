package routes

import (
	"github.com/cy77cc/hioshop/controllers/index"
	"github.com/gin-gonic/gin"
)

func RegisterIndexRoutes(router *gin.RouterGroup) {
	indexGroup := router.Group("/index")
	indexGroup.GET("/appInfo", index.AppInfo)
}
