package api

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/mauwahid/kafman/internal/app/publisher"
	"github.com/mauwahid/kafman/internal/platform/errs"
	"github.com/mauwahid/kafman/internal/presenter/dto"
	"net/http"
)

type PublisherHandler struct {
	pubApp *publisher.Publisher
}

func NewPublisherHandler() *PublisherHandler {
	pubApp := publisher.New()
	return &PublisherHandler{pubApp: pubApp}
}

func (p *PublisherHandler) Publish(c echo.Context) error {

	var (
		pubReq dto.PubRequest
		pubRes dto.PubResponse
		err    error
	)

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(c.Request().Body); err != nil {
		return c.JSON(http.StatusBadRequest, errs.RespBadRequest)
	}

	pubReq.Topic = c.Param("topic")
	pubReq.Message = buf.Bytes()

	if !pubReq.Validate() {
		return c.JSON(http.StatusBadRequest, errs.RespBadRequest)
	}

	if pubRes, err = p.pubApp.Publish(pubReq); err != nil {
		return c.JSON(http.StatusInternalServerError, pubRes)
	}

	return c.JSON(http.StatusOK, pubRes)
}
