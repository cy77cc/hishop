package goods

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGoods(c *gin.Context) {
	db := util.GetDB()
	keyword := c.Query("keyword")
	sort := c.Query("sort")
	order := c.Query("order")
	sales := c.Query("sales")

	goods := models.Goods{}
	if sort == "price" {
		db.Where("name like ?", "%"+keyword+"%").
			Order("retail_price " + order).
			Find(&goods)
	} else if sort == "sales" {
		db.Where("name like ?", "%"+keyword+"%").
			Order("sell_volume " + sales).
			Find(&goods)
	} else {
		db.Where("name like ?", "%"+keyword+"%").
			Order("sort_order asc").
			Find(&goods)
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   goods,
	})
}
