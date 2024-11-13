package goods

import (
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Count(c *gin.Context) {
	db := util.GetDB()
	var goodsCount int
	db.Table("hiolabs_goods").Where("is_delete=0").Select("count(*)").Find(&goodsCount)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"goodsCount": goodsCount,
		},
	})
}
