package cart

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetGoodsCount(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

}
