package catalog

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func CurrentList(c *gin.Context) {
	db := util.GetDB()
	req := struct {
		Id   int `json:"id"`
		Size int `json:"size"`
		Page int `json:"page"`
	}{}
	goods := make([]models.Goods, 0)
	_ = c.BindJSON(&req)

	var count int

	db.Where("category_id=?", req.Id).
		Order("sort_order").
		Offset(req.Size * (req.Page - 1)).Limit(req.Size).Find(&goods)
	db.Model(&models.Goods{}).
		Where("category_id=?", req.Id).
		Select("count(*) as count").Find(&count)
	totalPages := int(math.Ceil(float64(count) / float64(req.Size)))
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"count":       count,
			"totalPages":  totalPages,
			"pageSize":    req.Size,
			"currentPage": req.Page,
			"data":        goods,
		},
	})
}
