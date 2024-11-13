package catalog

import (
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CurrentCatalog(c *gin.Context) {
	db := util.GetDB()
	id, _ := strconv.Atoi(c.Query("id"))
	category := models.Category{}

	db.Where("id=?", id).Find(&category)

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   category,
	})
}
