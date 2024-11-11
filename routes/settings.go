package routes

import (
	"github.com/cy77cc/hioshop/controllers/settings"
	"github.com/gin-gonic/gin"
)

func RegisterSettingsRoutes(router *gin.RouterGroup) {
	settingsGroup := router.Group("/settings")
	settingsGroup.GET("/showSettings", settings.ShowSettings)
}
