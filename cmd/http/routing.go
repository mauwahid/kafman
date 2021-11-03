package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mauwahid/kafman/internal/presenter/api"
	"net/http"
)

func setupApiRoute(e *echo.Echo) {
	pub := api.NewPublisherHandler()
	e.POST("/kafman/v1/publish/:topic", pub.Publish)
	e.Any("/kafman/ping", echoPing)
}

func echoPing(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
