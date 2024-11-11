package settings

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowSettings(c *gin.Context) {
	db := util.DBOpenShop()
	setting := models.ShowSettings{}
	db.Where("id=?", 1).Find(&setting)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   setting,
	})
}
