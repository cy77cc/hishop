package routes

import (
	"github.com/cy77cc/hioshop/controllers/search"
	"github.com/gin-gonic/gin"
)

func RegisterSearchRoutes(router *gin.RouterGroup) {
	searchGroup := router.Group("/search")
	searchGroup.GET("/index", search.Index)
}
