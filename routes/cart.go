package routes

import (
	"github.com/cy77cc/hioshop/controllers/cart"
	"github.com/gin-gonic/gin"
)

func RegisterCartGroup(router *gin.RouterGroup) {
	cartGroup := router.Group("/cart")
	cartGroup.GET("/goodsCount", cart.GetGoodsCount)
}
