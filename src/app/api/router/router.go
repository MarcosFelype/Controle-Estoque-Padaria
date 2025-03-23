package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Start() *echo.Echo {
	server := echo.New()

	api := server.Group("/api")

	api.GET("/docs/*", echoSwagger.WrapHandler)

	loadProductRotes(api)
	return server
}
