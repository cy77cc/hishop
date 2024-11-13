package search

import (
	"github.com/cy77cc/hioshop/configs"
	"github.com/cy77cc/hioshop/middleware"
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	db := util.GetDB()
	// 初始化一个jwt对象，解析token
	auth := middleware.GetAuthMiddleware()
	jwt, _ := auth.GetClaimsFromJWT(c)
	openid := jwt[configs.IdentityKey]

	var keyword models.Keywords
	var user models.User
	historyKeywordList := make([]string, 0)
	hotKeywordList := make([]string, 0)
	db.Where("is_default=1 and is_show=1").First(&keyword)
	db.Table("hiolabs_keywords").Where("is_hot=1").Distinct("keyword").First(&hotKeywordList)
	db.Where("weixin_openid=?", openid).Find(&user)
	db.Table("hiolabs_search_history").
		Distinct("keyword").
		Where("user_id=?", user.ID).Find(&historyKeywordList)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"defaultKeyword":     keyword,
			"historyKeywordList": historyKeywordList,
			"hotKeywordList":     hotKeywordList,
		},
	})
}
