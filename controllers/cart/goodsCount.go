package cart

import (
	"github.com/cy77cc/hioshop/configs"
	"github.com/cy77cc/hioshop/middleware"
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGoodsCount 获取购物车商品数量
// @Summary 获取购物车商品数量
// @Description 获取购物车商品数量
// @Tags 购物车相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "hio 用户令牌"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/v1/cart/goodsCount
func GetGoodsCount(c *gin.Context) {
	auth := middleware.GetAuthMiddleware()
	jwt, _ := auth.GetClaimsFromJWT(c)
	openid := jwt[configs.IdentityKey]
	db := util.GetDB()

	var userId int
	var cartCount int
	db.Model(&models.User{}).Where("weixin_openid=?", openid).Select("id").Find(&userId)
	db.Model(&models.Cart{}).Where("user_id=?", userId).Group("user_id").Select("count(*)").Find(&cartCount)

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"cartTotal": map[string]interface{}{
				"goodsCount": cartCount,
			},
		},
	})
}
