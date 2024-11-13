package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"cartList": "",
			"cartTotal": map[string]interface{}{
				"goodsCount":         2,
				"goodsAmount":        "0.40",
				"checkedGoodsCount":  0,
				"checkedGoodsAmount": "0.00",
				"user_id":            5590,
				"numberChange":       0,
			},
		},
	})
}
