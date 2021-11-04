package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mauwahid/kafman/internal/presenter/api"
)

func setupApiRoute(e *echo.Echo) {
	pub := api.NewPublisherHandler()
	g := e.Group("/kafman")
	g.POST("/v1/publish/:topic", pub.Publish)
	g.GET("/ping", echoPing)
}

func echoPing(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
