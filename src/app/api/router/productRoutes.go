package router

import (
	dicontainter "padaria/src/app/api/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadProductRotes(api *echo.Group) {
	productGroup := api.Group("/product")

	productHandlers := dicontainter.GetProductHandlers()

	productGroup.POST("/new", productHandlers.PostProduct) //função dentro de outra função

	productGroup.GET("", productHandlers.GetProducts)
}
