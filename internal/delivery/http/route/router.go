package route

import (
	"boiler-plate-clean/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

type Router struct {
	App            *gin.Engine
	ProductHandler *http.ProductHTTPHandler
	StockHandler   *http.StockHTTPHandler
	SaleHandler    *http.SaleHTTPHandler
}

func (h *Router) Setup() {
	api := h.App.Group("")
	{

		//Product Routes
		productApi := api.Group("/product")
		//campaignApi.Use(h.RequestMiddleware.RequestHeader)
		{
			productApi.POST("", h.ProductHandler.Create)
			productApi.GET("", h.ProductHandler.Find)
			productApi.GET("/:id", h.ProductHandler.FindOne)
			productApi.PUT("/:id", h.ProductHandler.Update)
			productApi.DELETE("/:id", h.ProductHandler.Delete)
		}
		stockApi := api.Group("/stock")
		//campaignApi.Use(h.RequestMiddleware.RequestHeader)
		{
			stockApi.POST("", h.StockHandler.Create)
			stockApi.GET("", h.StockHandler.Find)
			stockApi.GET("/:id", h.StockHandler.FindOne)
			stockApi.PUT("/:id", h.StockHandler.Update)
			stockApi.DELETE("/:id", h.StockHandler.Delete)
		}
		saleApi := api.Group("/sale")
		//campaignApi.Use(h.RequestMiddleware.RequestHeader)
		{
			saleApi.POST("", h.SaleHandler.Create)
			saleApi.GET("/report", h.SaleHandler.GetProfitReport)
			//saleApi.GET("", h.SaleHandler.Find)
			//saleApi.GET("/:id", h.SaleHandler.FindOne)
			//saleApi.PUT("/:id", h.SaleHandler.Update)
			//saleApi.DELETE("/:id", h.SaleHandler.Delete)
		}
	}
}
