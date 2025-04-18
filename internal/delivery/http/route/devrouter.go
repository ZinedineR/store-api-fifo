package route

import (
	docs "boiler-plate-clean/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Router) setupDevRouter() {

}

func (h *Router) SwaggerRouter() {
	docs.SwaggerInfo.BasePath = ""

	//h.router.Use(static.Serve("/"+h.base.AppConfig.AppConfig.AppName+"/api/"+h.base.AppConfig.AppConfig.AppVersion+"/static", static.LocalFile("./docs", true)))
	//h.router.Static("/"+h.base.AppConfig.AppConfig.AppName+"/api/"+h.base.AppConfig.AppConfig.AppVersion+"/static/*any", "./docs")
	h.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
