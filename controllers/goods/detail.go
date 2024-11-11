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
	GoodsNumber int `json:"goods_number"`
}

func Detail(c *gin.Context) {
	goods_id, _ := strconv.Atoi(c.Query("id"))
	db := util.DBOpenShop()
	info := models.Goods{}
	gallery := make([]models.GoodsGallery, 0)
	product := make([]models.Product, 0)
	specification := models.Specification{}
	value := make([]Value, 0)
	db.Where("id=?", goods_id).Find(&info)
	db.Where("goods_id=?", goods_id).Find(&gallery)
	db.Where("goods_id=?", goods_id).Find(&product)
	db.Table("hiolabs_goods_specification").
		Where("goods_id=? and is_delete=0", goods_id).
		Find(&value)
	for k, v := range value {
		goodsNumber := 0
		db.Table("hiolabs_product").
			Where("goods_specification_ids=?", v.Id).
			Select("goods_number").Find(&goodsNumber)
		value[k].GoodsNumber = goodsNumber
	}
	db.Table("hiolabs_specification").
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
				"specification_id": specification.Id,
				"name":             specification.Name,
				"valueList":        value,
			},
		},
	})
}
