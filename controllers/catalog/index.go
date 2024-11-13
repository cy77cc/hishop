package catalog

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	db := util.GetDB()
	categoryList := make([]models.Category, 0)
	db.Where("is_show=1").Order("sort_order").Find(&categoryList)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"categoryList": categoryList,
		},
	})
}
