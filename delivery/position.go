package delivery

import (
	"net/http"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type positionDelivery struct {
	positionUsecase model.PositionUsecase
}

type PositionDelivery interface {
	Mount(group *echo.Group)
}

// NewPositionDelivery returns the delivery implementation of the position group path
func NewPositionDelivery(positionUsecase model.PositionUsecase) PositionDelivery {
	return &positionDelivery{positionUsecase: positionUsecase}
}

// Mount mounts the available paths of the position group path
func (p *positionDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchPositionHandler)
	group.POST("", p.StorePositionHandler)
	group.GET("/:id", p.DetailPositionHandler)
	group.DELETE("/:id", p.DeletePositionHandler)
	group.PUT("/:id", p.EditPositionHandler)
}

// FetchPositionHandler handles the delivery of the path that gets all position data
func (p *positionDelivery) FetchPositionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	positionList, err := p.positionUsecase.FetchPosition(ctx, limitInt, offsetInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "success", positionList)

}

// StorePositionHandler handles the delivery of the path that inserts a new position data
func (p *positionDelivery) StorePositionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.PositionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	position, err := p.positionUsecase.StorePosition(ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return helper.ResponseSuccessJson(c, "success", position)
}

// DetailPositionHandler handles the delivery of the path that gets position data
func (p *positionDelivery) DetailPositionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	position, err := p.positionUsecase.GetByID(ctx, IdInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", position)

}

// DeletePositionHandler handles the delivery of the path that deletes position data
func (p *positionDelivery) DeletePositionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	err := p.positionUsecase.DestroyPosition(ctx, IdInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", "")

}

// EditPositionHandler handles the delivery of the path that updates position data
func (p *positionDelivery) EditPositionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.PositionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	id := c.Param("id")
	IdInt, _ := strconv.Atoi(id)

	position, err := p.positionUsecase.EditPosition(ctx, IdInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success edit", position)
}
