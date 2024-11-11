package routes

import (
	"github.com/cy77cc/hioshop/controllers/goods"
	"github.com/gin-gonic/gin"
)

func RegisterGoodsRoutes(router *gin.RouterGroup) {
	goodsGroup := router.Group("/goods")
	goodsGroup.GET("/detail", goods.Detail)
}
