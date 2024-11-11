package main

import (
	"github.com/cy77cc/hioshop/routes"
	"log"
	"net/http"

	"github.com/cy77cc/hioshop/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// the jwt middleware
	authMiddleware := middleware.GetAuthMiddleware()

	//register middleware
	engine.Use(middleware.HandlerMiddleWare(authMiddleware), middleware.Cors())
	//register route
	middleware.RegisterRoute(engine, authMiddleware)

	engine.Static("/static", "./www/static")
	v1 := engine.Group("/api/v1")

	routes.RegisterSettingsRoutes(v1)
	routes.RegisterSearchRoutes(v1)
	routes.RegisterGoodsRoutes(v1)
	routes.RegisterIndexRoutes(v1)
	routes.RegisterCartGroup(v1)

	// start http server
	if err := http.ListenAndServe(":8080", engine); err != nil {
		log.Fatal(err)
	}
}
