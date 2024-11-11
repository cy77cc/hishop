package cart

import (
	"github.com/cy77cc/hioshop/configs"
	"github.com/cy77cc/hioshop/middleware"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGoodsCount(c *gin.Context) {
	auth := middleware.GetAuthMiddleware()
	jwt, _ := auth.GetClaimsFromJWT(c)
	openid := jwt[configs.IdentityKey]
	db := util.GetDB()

	var userId int
	var cartCount int
	db.Table("hiolabs_user").Where("weixin_openid=?", openid).Select("id").Find(&userId)
	db.Table("hiolabs_cart").Where("user_id=?", userId).Group("user_id").Select("count(*)").Find(&cartCount)

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
