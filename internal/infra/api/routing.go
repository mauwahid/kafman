package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mauwahid/kafman/internal/interfaces/api"
	"net/http"
)

func SetupApiRoute(e *echo.Echo) {

	pub := api.NewPublisherHandler()
	e.POST("/kafman/v1/publish", pub.Publish)
	e.Any("/kafman/ping", echoPing)
}

func echoPing(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
