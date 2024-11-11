package index

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cy77cc/hioshop/configs"
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type banner struct {
	LinkType int    `json:"link_type"`
	GoodsId  int    `json:"goods_id"`
	ImageUrl string `json:"image_url"`
	Link     string `json:"link"`
}

type categoryListItem struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Banner    string         `json:"banner"`
	Height    int            `json:"height"`
	GoodsList []models.Goods `json:"goodsList" gorm:"foreignKey:CategoryId;reference:Id"`
}

func AppInfo(c *gin.Context) {

	db := util.DBOpenShop()
	channels := make([]models.Category, 0)
	db.Where("is_channel=? and parent_id=?", 1, 0).Order("sort_order").Find(&channels)
	banners := make([]banner, 0)
	db.Table("hiolabs_ad").Where("enabled=? and is_delete=?", 1, 0).Find(&banners)
	notices := make([]models.Notice, 0)
	db.Where("is_delete=?", 0).Find(&notices)

	categoryList := make([]categoryListItem, 0)
	db.Table("hiolabs_category").
		Select("id", "name", "img_url as banner", " p_height as height").
		Where("parent_id=? and is_show=?", 0, 1).
		Order("sort_order").
		Find(&categoryList)

	for k, v := range categoryList {
		goods := make([]models.Goods, 0)
		db.Table("hiolabs_goods").Where(map[string]interface{}{
			"category_id": v.Id,
			"is_on_sale":  1,
			"is_index":    1,
			"is_delete":   0,
		}).Where("goods_number>=?", 0).Order("sort_order").Find(&goods)
		categoryList[k].GoodsList = goods
	}

	claims := jwt.ExtractClaims(c)
	openid, _ := claims[configs.IdentityKey].(string)

	var userId int
	var cartCount int
	db.Table("hiolabs_user").Where("open_id=?", openid).Select("id").Find(&userId)
	db.Table("hiolabs_cart").Where("user_id=?", userId).Group("user_id").Select("count(*)").Find(&cartCount)

	data := make(map[string]interface{})
	data["channel"] = channels
	data["banner"] = banners
	data["notice"] = notices
	data["categoryList"] = categoryList
	data["cartCount"] = cartCount
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   data,
	})
}
