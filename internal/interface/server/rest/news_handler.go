package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/cast"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	serviceNews "github.com/mfsyahrz/bareksaproject/internal/service/news"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type NewsHandler struct {
	service serviceNews.NewsService
}

func NewNewsHandler(service serviceNews.NewsService) *NewsHandler {
	return &NewsHandler{service}
}

func (h *NewsHandler) FindAll(c echo.Context) error {
	var (
		resp    = Response{}
		status  = constant.NewsStatus(c.QueryParam("status"))
		TopicID = cast.ToInt64(c.QueryParam("topic-id"))
	)

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	data, err := h.service.FindAll(c.Request().Context(), status, TopicID)
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess
	resp.Data = data

	return nil
}

func (h *NewsHandler) Create(c echo.Context) error {
	var (
		resp   = Response{}
		entity = entity.News{}
	)

	if err := c.Bind(&entity); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(entity); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	err := h.service.Create(c.Request().Context(), entity)
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess

	return nil
}

func (h *NewsHandler) UpdateByID(c echo.Context) error {
	var (
		resp   = Response{}
		id     = cast.ToInt64(c.Param("id"))
		entity = entity.News{}
	)

	if id == 0 {
		resp.Message = constant.IDEmpty
		return c.JSON(http.StatusBadRequest, resp)
	}

	entity.ID = id
	if err := c.Bind(&entity); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(entity); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	err := h.service.UpdateByID(c.Request().Context(), entity)
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess

	return nil
}

func (h *NewsHandler) DeleteByID(c echo.Context) error {
	var (
		resp = Response{}
		id   = cast.ToInt64(c.Param("id"))
	)

	if id == 0 {
		resp.Message = constant.IDEmpty
		return c.JSON(http.StatusBadRequest, resp)
	}

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	err := h.service.DeleteByID(c.Request().Context(), id)
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess

	return nil
}

func (h *NewsHandler) FindOne(c echo.Context) error {
	var (
		resp = Response{}
		id   = cast.ToInt64(c.Param("id"))
	)

	if id == 0 {
		resp.Message = constant.IDEmpty
		return c.JSON(http.StatusBadRequest, resp)
	}

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	data, err := h.service.FindOne(c.Request().Context(), id)
	if err != nil {
		resp.Message = err.Error()
		return nil
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess
	resp.Data = data

	return nil
}
