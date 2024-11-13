package routes

import (
	"github.com/cy77cc/hioshop/controllers/catalog"
	"github.com/gin-gonic/gin"
)

func RegisterCatalogGroup(router *gin.RouterGroup) {
	catalogGroup := router.Group("/catalog")
	catalogGroup.GET("/index", catalog.Index)
	catalogGroup.POST("/currentlist", catalog.CurrentList)
	catalogGroup.GET("/current", catalog.CurrentCatalog)
}
