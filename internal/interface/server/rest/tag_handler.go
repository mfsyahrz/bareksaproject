package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/cast"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	serviceTag "github.com/mfsyahrz/bareksaproject/internal/service/tag"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type TagHandler struct {
	service serviceTag.TagService
}

func NewTagHandler(service serviceTag.TagService) *TagHandler {
	return &TagHandler{service}
}

func (h *TagHandler) FindAll(c echo.Context) error {
	var resp Response

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	data, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		resp.Message = err.Error()
		return err
	}

	resp.Success = true
	resp.Message = constant.ResponseSuccess
	resp.Data = data

	return nil
}

func (h *TagHandler) FindOne(c echo.Context) error {
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

func (h *TagHandler) Create(c echo.Context) error {
	var (
		resp   = Response{}
		entity = entity.Tag{}
	)

	if err := c.Bind(&entity); err != nil {
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

func (h *TagHandler) UpdateByID(c echo.Context) error {
	var (
		resp   = Response{}
		id     = cast.ToInt64(c.Param("id"))
		entity = entity.Tag{}
	)

	if id == 0 {
		resp.Message = constant.IDEmpty
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Bind(&entity); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	entity.ID = id

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

func (h *TagHandler) DeleteByID(c echo.Context) error {
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
