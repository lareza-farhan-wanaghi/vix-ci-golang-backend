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

type userDelivery struct {
	userUsecase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

// NewUserDelivery returns the delivery implementation of the user group path
func NewUserDelivery(userUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{userUsecase: userUsecase}
}

// Mount mounts the available paths of the user group path
func (p *userDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchUserHandler)
	group.POST("", p.StoreUserHandler)
	group.GET("/:id", p.DetailUserHandler)
	group.DELETE("/:id", p.DeleteUserHandler)
	group.PUT("/:id", p.EditUserHandler)
	group.POST("/withdraw", p.WithdrawHandler)
}

// FetchUserHandler handles the delivery of the path that gets all user data
func (p *userDelivery) FetchUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	userList, err := p.userUsecase.FetchUser(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return helper.ResponseSuccessJson(c, "success", userList)

}

// StoreUserHandler handles the delivery of the path that inserts user data
func (p *userDelivery) StoreUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())

	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	user, err := p.userUsecase.StoreUser(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "success", user)
}

// DetailUserHandler handles the delivery of the path that gets user data
func (p *userDelivery) DetailUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	user, err := p.userUsecase.GetByID(ctx, IdInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "", user)

}

// DeleteUserHandler handles the delivery of the path that deletes user data
func (p *userDelivery) DeleteUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	err := p.userUsecase.DestroyUser(ctx, IdInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", "")

}

// EditUserHandler handles the delivery of the path that updates user data
func (p *userDelivery) EditUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())

	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	id := c.Param("id")
	IdInt, _ := strconv.Atoi(id)

	user, err := p.userUsecase.EditUser(ctx, IdInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success edit", user)
}

// WithdrawHandler handles the delivery of the path that withdraws a user's salary
func (p *userDelivery) WithdrawHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.WithdrawRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	err := p.userUsecase.WithdrawSalary(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success withdraw salary", "")

}
