package main

import (
	docs "github.com/cy77cc/hioshop/docs"
	"github.com/cy77cc/hioshop/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"github.com/cy77cc/hioshop/middleware"
	"github.com/gin-gonic/gin"
)

// @title hioshop
// @version 1.0
// @description 测试
// @termsOfService http://swagger.io/terms/

// @contact.name cy77cc
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 8080
// @BasePath /api/v1
func main() {
	engine := gin.Default()
	// the jwt middleware
	authMiddleware := middleware.GetAuthMiddleware()

	//register middleware
	engine.Use(middleware.HandlerMiddleWare(authMiddleware), middleware.Cors())
	//register route
	middleware.RegisterRoute(engine, authMiddleware)

	engine.Static("/static", "./www/static")
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := engine.Group("/api/v1")

	routes.RegisterSettingsRoutes(v1)
	routes.RegisterSearchRoutes(v1)
	routes.RegisterGoodsRoutes(v1)
	routes.RegisterIndexRoutes(v1)
	routes.RegisterCartGroup(v1)
	routes.RegisterCatalogGroup(v1)

	// start http server
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := http.ListenAndServe(":8080", engine); err != nil {
		log.Fatal(err)
	}
}
