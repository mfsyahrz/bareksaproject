package rest

import (
	"net/http"

	"github.com/labstack/echo"

	serviceTopic "github.com/mfsyahrz/bareksaproject/internal/service/topic"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type TopicHandler struct {
	service serviceTopic.TopicService
}

func NewTopicHandler(service serviceTopic.TopicService) *TopicHandler {
	return &TopicHandler{service}
}

func (h *TopicHandler) FindAll(c echo.Context) error {
	var resp Response

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	data, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess
	resp.Data = data

	return nil
}
