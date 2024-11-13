package goods

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Value struct {
	models.GoodsSpecification
	GoodsNumber int `json:"goods_number" gorm:"-"`
}

func Detail(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("id"))
	db := util.GetDB()
	info := models.Goods{}
	gallery := make([]models.GoodsGallery, 0)
	product := make([]models.Product, 0)
	specification := models.Specification{}
	value := make([]Value, 0)
	db.Where("id=?", goodsId).Find(&info)
	db.Where("goods_id=?", goodsId).Find(&gallery)
	db.Where("goods_id=?", goodsId).Find(&product)

	db.Model(&models.GoodsSpecification{}).
		Where("goods_id=? and is_delete=0", goodsId).
		Find(&value)

	for k, v := range value {
		goodsNumber := 0
		db.Model(&models.Product{}).
			Where("goods_specification_ids=?", v.ID).
			Select("goods_number").Find(&goodsNumber)
		value[k].GoodsNumber = goodsNumber
	}
	db.Model(&models.Specification{}).
		Where("id=?", value[0].SpecificationId).
		Find(&specification)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"info":    info,
			"gallery": gallery,
			"product": product,
			"specificationList": map[string]interface{}{
				"specification_id": specification.ID,
				"name":             specification.Name,
				"valueList":        value,
			},
		},
	})
}
